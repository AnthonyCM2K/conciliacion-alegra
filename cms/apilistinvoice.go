package cms

import (
	"cmsalegra/configuration"
	"cmsalegra/model"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
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
	var totalInvoicesCMS int
	var totalAmountCMS float64

	// Crear un mapa para almacenar los IDs de facturas procesadas
	processedInvoiceIDs := make(map[string]bool)

	for _, invoice := range invoices.Data {
		// Convertir el ID de la factura a string
		invoiceIDStr := strconv.FormatInt(invoice.ID, 10)

		// Verificar si la factura ya ha sido procesada
		if _, exists := processedInvoiceIDs[invoiceIDStr]; !exists {
			// Si no ha sido procesada, agregarla al mapa
			processedInvoiceIDs[invoiceIDStr] = true

			var invoiceItems []model.InvoiceItemListResponse

			if invoice.AlegraTransactionList != nil && invoice.AlegraTransactionList.InvoiceRelationList.InvoiceItems != nil {
				for _, itemPrice := range invoice.AlegraTransactionList.InvoiceRelationList.InvoiceItems {
					invoiceItem := model.InvoiceItemListResponse{
						OriginalPrice: itemPrice.OriginalPrice,
					}
					invoiceItems = append(invoiceItems, invoiceItem)
					// Sumar el OriginalPrice de este item al total general
					totalAmountCMS += math.Round((itemPrice.OriginalPrice/invoice.ExchangeRate)*100) / 100
				}
			}

			invoiceRelation := model.InvoiceRelationListResponse{
				InvoiceItems: invoiceItems,
			}
			invoiceBanckAccount := model.AlegraDataListResponse{
				BankAccount: invoice.AlegraTransactionList.AlegraDataList.BankAccount,
			}

			alegraTransaction := model.AlegraTransactionListResponse{
				InvoiceRelationListResponse: invoiceRelation,
				AlegraPaymentID:             invoice.AlegraTransactionList.AlegraPaymentID,
				AlegraDataList:              invoiceBanckAccount,
			}

			newInvoice := model.InvoiceListResponse{
				ID:            invoice.ID,
				UserID:        int64(invoice.UserID),
				Email:         invoice.Email,
				FirstName:     invoice.FirstName,
				LastName:      invoice.LastName,
				ExchangeRate:  invoice.ExchangeRate,
				Currency:      invoice.Currency,
				PaymentMethod: invoice.PaymentMethod,

				AlegraTransactionListResponse: alegraTransaction,
			}

			resultInvoices = append(resultInvoices, newInvoice)

			// Incrementar el contador de facturas
			totalInvoicesCMS++
		}
	}

	// for _, invoice := range invoices.Data {

	// 	var invoiceItems []model.InvoiceItemListResponse

	// 	if invoice.AlegraTransactionList != nil && invoice.AlegraTransactionList.InvoiceRelationList.InvoiceItems != nil {
	// 		for _, itemPrice := range invoice.AlegraTransactionList.InvoiceRelationList.InvoiceItems {
	// 			invoiceItem := model.InvoiceItemListResponse{
	// 				OriginalPrice: itemPrice.OriginalPrice,
	// 			}
	// 			invoiceItems = append(invoiceItems, invoiceItem)
	// 			// Sumar el OriginalPrice de este item al total general
	// 			totalAmountCMS += math.Round((itemPrice.OriginalPrice/invoice.ExchangeRate)*100) / 100
	// 		}
	// 	}

	// 	invoiceRelation := model.InvoiceRelationListResponse{
	// 		InvoiceItems: invoiceItems,
	// 	}
	// 	invoiceBanckAccount := model.AlegraDataListResponse{
	// 		BankAccount: invoice.AlegraTransactionList.AlegraDataList.BankAccount,
	// 	}

	// 	alegraTransaction := model.AlegraTransactionListResponse{
	// 		InvoiceRelationListResponse: invoiceRelation,
	// 		AlegraPaymentID:             invoice.AlegraTransactionList.AlegraPaymentID,
	// 		AlegraDataList:              invoiceBanckAccount,
	// 	}

	// 	newInvoice := model.InvoiceListResponse{
	// 		ID:            invoice.ID,
	// 		UserID:        int64(invoice.UserID),
	// 		Email:         invoice.Email,
	// 		FirstName:     invoice.FirstName,
	// 		LastName:      invoice.LastName,
	// 		ExchangeRate:  invoice.ExchangeRate,
	// 		Currency:      invoice.Currency,
	// 		PaymentMethod: invoice.PaymentMethod,

	// 		AlegraTransactionListResponse: alegraTransaction,
	// 	}

	// 	resultInvoices = append(resultInvoices, newInvoice)

	// 	// Incrementar el contador de facturas
	// 	totalInvoicesCMS++
	// }

	jsonData, err := json.Marshal(resultInvoices)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("error marshalling result invoices: %w", err)
	}
	//fmt.Println(string(jsonData))
	return jsonData, totalInvoicesCMS, totalAmountCMS, nil

}
