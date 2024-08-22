package conciliacion

import (
	"cmsalegra/alegra"
	"cmsalegra/cms"
	"cmsalegra/configuration"
	"cmsalegra/model"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Conciliation ejecuta la conciliación de facturas de CMS con facturas de Alegra en un rango de fechas determinado.
func Conciliation(fecha string, config configuration.Configuration) {
	facturasAlegra, err := alegra.QueryApiByteAlegraBankAll(fecha, config)
	if err != nil {
		log.Fatal("consultaApiAlegra():", err)
	}

	facturas, err := cms.QueryApiByteCMSReports(fecha, config)
	if err != nil {
		log.Fatal("consultaApiCMS():", err)
	}

	cmsInvoices, alegraInvoices, err := unmarshalInvoices(facturas, facturasAlegra)
	if err != nil {
		log.Fatal(err)
	}

	cmsMap := createCMSMap(cmsInvoices)
	alegraMap := createAlegraMap(alegraInvoices)

	// Aquí corregimos los tipos para que coincidan correctamente
	notInCMS, notInAlegra := findMissingInvoices(cmsMap, alegraMap)

	NotPriceAlegra, NotPriceCMS := findPriceDiscrepancies(cmsMap, alegraInvoices)

	fileName := fecha + "_not_in_CMS_or_Alegra.csv"
	exportToCSV(fileName, notInCMS, notInAlegra)

	fileNameAmount := fecha + "_discrepancies_in_amount.csv"
	if err := exportToCSVAmount(fileNameAmount, NotPriceCMS, NotPriceAlegra); err != nil {
		log.Fatalf("Error exporting to CSV: %v", err)
	}

	fmt.Printf("CSV generado exitosamente: %s\n", fileNameAmount)
}

// findMissingInvoices Cambiamos los tipos de los mapas para que sean consistentes
func findMissingInvoices(cmsMap map[any]model.InvoiceListResponse, alegraMap map[any]model.InvoiceAlegraResponse) ([]model.InvoiceAlegraResponse, []model.InvoiceListResponse) {
	notInCMS := make([]model.InvoiceAlegraResponse, 0)
	notInAlegra := make([]model.InvoiceListResponse, 0)

	for id, alegra := range alegraMap {
		if _, ok := cmsMap[id]; !ok {
			notInCMS = append(notInCMS, alegra)
		}
	}

	for id, cms := range cmsMap {
		if _, ok := alegraMap[id]; !ok {
			notInAlegra = append(notInAlegra, cms)
		}
	}

	return notInCMS, notInAlegra
}

// unmarshalInvoices Decodificamos los archivos JSON
func unmarshalInvoices(facturasCMS, facturasAlegra []byte) ([]model.InvoiceListResponse, []model.InvoiceAlegraResponse, error) {
	var cmsInvoices []model.InvoiceListResponse
	if err := json.Unmarshal(facturasCMS, &cmsInvoices); err != nil {
		return nil, nil, fmt.Errorf("error decoding CMS invoices JSON: %v", err)
	}

	var alegraInvoices []model.InvoiceAlegraResponse
	if err := json.Unmarshal(facturasAlegra, &alegraInvoices); err != nil {
		return nil, nil, fmt.Errorf("error decoding Alegra invoices JSON: %v", err)
	}

	return cmsInvoices, alegraInvoices, nil
}

func createCMSMap(cmsInvoices []model.InvoiceListResponse) map[any]model.InvoiceListResponse {
	cmsMap := make(map[any]model.InvoiceListResponse, len(cmsInvoices))
	for _, invoice := range cmsInvoices {
		cmsMap[invoice.AlegraTransactionListResponse.AlegraPaymentID] = invoice
	}
	return cmsMap
}

func createAlegraMap(alegraInvoices []model.InvoiceAlegraResponse) map[any]model.InvoiceAlegraResponse {
	alegraMap := make(map[any]model.InvoiceAlegraResponse, len(alegraInvoices))
	for _, invoice := range alegraInvoices {
		key := strconv.FormatInt(invoice.ID, 10)
		alegraMap[key] = invoice
	}
	return alegraMap
}

// findPriceDiscrepancies Busca las facturas con discrepancias en el valor de la venta
func findPriceDiscrepancies(cmsMap map[any]model.InvoiceListResponse, alegraInvoices []model.InvoiceAlegraResponse) ([]model.InvoiceAlegraResponse, []model.InvoiceListResponse) {
	NotPriceAlegra := make([]model.InvoiceAlegraResponse, 0)
	NotPriceCMS := make([]model.InvoiceListResponse, 0)

	for _, alegra := range alegraInvoices {
		cms, ok := cmsMap[strconv.FormatInt(alegra.ID, 10)]
		if !ok {
			continue
		}

		totalItemsAmount := totalCMSInvoiceAmount(cms)

		if totalItemsAmount != alegra.Amount {
			NotPriceAlegra = append(NotPriceAlegra, alegra)
			NotPriceCMS = append(NotPriceCMS, cms)
		}
	}

	return NotPriceAlegra, NotPriceCMS
}

// exportToCSV Exporta a un archivo CSV las facturas que están registradas en un solo sistema, ya sea en CMS o en Alegra, pero no en ambos.
func exportToCSV(filename string, notInCMS []model.InvoiceAlegraResponse, notInAlegra []model.InvoiceListResponse) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Error al crear el archivo CSV: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Escribir el encabezado
	err = writer.Write([]string{"Source", "ID", "UserID", "Email User", "AlegraPaymentID", "Amount", "Detail", "Total Invoices", "Total Amount"})
	if err != nil {
		log.Fatalf("Error al escribir el encabezado CSV: %v", err)
	}

	var totalAlegra, totalCMS int
	var sumAlegra, sumCMS float64

	// Escribir datos de `notInCMS` (facturas en Alegra no presentes en CMS)
	for _, invoice := range notInCMS {
		err := writer.Write([]string{
			"Invoice Alegra not in CMS",
			strconv.FormatInt(invoice.ID, 10),
			"", // No hay `UserID` en `InvoiceAlegraResponse`
			"",
			"", // No hay `AlegraPaymentID` en `InvoiceAlegraResponse`
			fmt.Sprintf("%.2f", invoice.Amount),
			"", "", "",
		})
		if err != nil {
			log.Fatalf("Error al escribir datos de notInCMS: %v", err)
		}
		totalAlegra++
		sumAlegra += invoice.Amount
	}

	// Escribir datos de `notInAlegra` (facturas en CMS no presentes en Alegra)
	for _, invoice := range notInAlegra {
		totalAmount := totalCMSInvoiceAmount(invoice)
		err := writer.Write([]string{
			"Invoice CMS not in Alegra",
			strconv.FormatInt(invoice.ID, 10),
			strconv.FormatInt(invoice.UserID, 10),
			invoice.Email,
			invoice.AlegraTransactionListResponse.AlegraPaymentID,
			fmt.Sprintf("%.2f", totalAmount), // No hay `Amount` en `InvoiceListResponse` pero calculamos suma Originalprice
			"", "", "",
		})
		if err != nil {
			log.Fatalf("Error al escribir datos de notInAlegra: %v", err)
		}
		totalCMS++
		sumCMS += totalAmount

	}

	// Escribir los totales
	err = writer.Write([]string{
		"", "", "", "", "", "", "Total facturas Alegra:",
		strconv.Itoa(totalAlegra) + " registros",
		fmt.Sprintf("%.2f", sumAlegra),
	})
	if err != nil {
		log.Fatalf("Error al escribir el total de Alegra: %v", err)
	}

	err = writer.Write([]string{
		"", "", "", "", "", "", "Total facturas CMS:",
		strconv.Itoa(totalCMS) + " registros",
		fmt.Sprintf("%.2f", sumCMS),
	})
	if err != nil {
		log.Fatalf("Error al escribir el total de CMS: %v", err)
	}

	fmt.Printf("Archivo CSV generado correctamente: %s\n", filename)
}

// exportToCSVAmount exporta las facturas con discrepancias
func exportToCSVAmount(fileName string, notPriceCMS []model.InvoiceListResponse, notPriceAlegra []model.InvoiceAlegraResponse) error {

	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("error creating CSV file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Escribir el encabezado
	header := []string{"AlegraPaymentID", "CMS Invoice ID", "CMS Invoice Amount", "Alegra Invoice Amount", "Discrepancy Count"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("error writing header to CSV: %v", err)
	}

	// Escribir las facturas con discrepancias
	for i := range notPriceCMS {
		cms := notPriceCMS[i]
		alegra := notPriceAlegra[i]

		record := []string{
			cms.AlegraTransactionListResponse.AlegraPaymentID,
			strconv.FormatInt(cms.ID, 10),
			fmt.Sprintf("%.2f", totalCMSInvoiceAmount(cms)),
			fmt.Sprintf("%.2f", alegra.Amount),
		}

		if err := writer.Write(record); err != nil {
			return fmt.Errorf("error writing record to CSV: %v", err)
		}

	}
	// Escribir el total de discrepancias al final del archivo
	totalDiscrepancies := len(notPriceCMS)
	totalRecord := []string{
		"", "", "", "", "Total Discrepancies:", strconv.Itoa(totalDiscrepancies),
	}

	if err := writer.Write(totalRecord); err != nil {
		return fmt.Errorf("error writing total discrepancies to CSV: %v", err)
	}

	return nil
}

func totalCMSInvoiceAmount(cms model.InvoiceListResponse) float64 {
	var total float64
	for _, item := range cms.AlegraTransactionListResponse.InvoiceRelationListResponse.InvoiceItems {
		total += float64(item.OriginalPrice)
	}
	return total
}
