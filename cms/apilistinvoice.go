package cms

import (
	"cmsalegra/configuration"
	"cmsalegra/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// requestCMS consulta al endpoint de CMS
func requestCMS(fecha string, config configuration.Configuration) (*http.Response, error) {

	client := &http.Client{}

	//Peticion de token para CMS
	tokenCMS, err := ApiCMSLogin(config)
	if err != nil {
		log.Fatal(err)
	}

	// Base URL
	baseURL := config.CMSApi.UrlApiCmsConsulta

	// Parámetros de consulta
	params := url.Values{}
	params.Add("begins", fecha)
	params.Add("ends", fecha)

	// Construir la URL completa con parámetros
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+tokenCMS)

	response, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	return response, nil
}
func QueryApiByteCMSReports(fecha string, config configuration.Configuration) ([]byte, int, float64, error) {

	response, err := requestCMS(fecha, config)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var invoices model.InvoiceCMSList
	if err := json.NewDecoder(response.Body).Decode(&invoices); err != nil {
		return nil, 0, 0, fmt.Errorf("error decoding response: %w", err)
	}

	var resultInvoices []model.InvoiceListResponse
	invoiceMap := make(map[int64]model.InvoiceListResponse) // Definir el mapa fuera del bucle
	var totalInvoicesCMS int
	var totalAmountCMS float64

	for _, invoice := range invoices.Data {
		if existingInvoice, exists := invoiceMap[invoice.ID]; exists {
			// Si la factura ya existe, sumar InUsd y combinar los items
			existingInvoice.InUsd += invoice.InUsd
			totalAmountCMS += invoice.InUsd

			if invoice.AlegraTransactionList != nil && invoice.AlegraTransactionList.InvoiceRelationList.InvoiceItems != nil {
				for _, item4 := range invoice.AlegraTransactionList.InvoiceRelationList.InvoiceItems {
					invoiceItem := model.InvoiceItemListResponse{
						OriginalPrice: item4.OriginalPrice,
					}
					existingInvoice.AlegraTransactionListResponse.InvoiceRelationListResponse.InvoiceItems = append(existingInvoice.AlegraTransactionListResponse.InvoiceRelationListResponse.InvoiceItems, invoiceItem)
				}
			}
			invoiceMap[invoice.ID] = existingInvoice
		} else {
			// Si es una nueva factura, crearla
			var invoiceItems []model.InvoiceItemListResponse

			if invoice.AlegraTransactionList != nil && invoice.AlegraTransactionList.InvoiceRelationList.InvoiceItems != nil {
				for _, item4 := range invoice.AlegraTransactionList.InvoiceRelationList.InvoiceItems {
					invoiceItem := model.InvoiceItemListResponse{
						OriginalPrice: item4.OriginalPrice,
					}
					invoiceItems = append(invoiceItems, invoiceItem)
				}
			}

			invoiceRelation := model.InvoiceRelationListResponse{
				InvoiceItems: invoiceItems,
			}

			alegraTransaction := model.AlegraTransactionListResponse{
				InvoiceRelationListResponse: invoiceRelation,
				AlegraPaymentID:             invoice.AlegraTransactionList.AlegraPaymentID,
			}

			newInvoice := model.InvoiceListResponse{
				ID:           invoice.ID,
				UserID:       int64(invoice.UserID),
				Email:        invoice.Email,
				FirstName:    invoice.FirstName,
				LastName:     invoice.LastName,
				InUsd:        invoice.InUsd,
				ExchangeRate: invoice.ExchangeRate,

				AlegraTransactionListResponse: alegraTransaction,
			}

			invoiceMap[invoice.ID] = newInvoice
			totalAmountCMS += invoice.InUsd
		}
	}

	for _, invoice := range invoiceMap {
		resultInvoices = append(resultInvoices, invoice)
		totalInvoicesCMS++
	}

	jsonData, err := json.Marshal(resultInvoices)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("error marshalling result invoices: %w", err)
	}

	return jsonData, totalInvoicesCMS, totalAmountCMS, nil
}
