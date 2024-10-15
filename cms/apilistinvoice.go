package cms

import (
	"cmsalegra/configuration"
	"cmsalegra/model"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
)

func RequestCMS(fecha string, config configuration.Configuration, client CMSClient) (*http.Response, error) {
	token, err := client.ApiCMSLogin(config)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", config.CMSApi.UrlApiCmsConsulta, nil)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Add("begins", fecha)
	query.Add("ends", fecha)
	req.URL.RawQuery = query.Encode()

	req.Header.Set("Authorization", "Bearer "+token)

	clientHttp := &http.Client{}
	return clientHttp.Do(req)
}

func QueryApiByteCMSReports(fecha string, config configuration.Configuration, client CMSClient) ([]byte, int, float64, error) {
	response, err := RequestCMS(fecha, config, client)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var invoices model.InvoiceCMSList
	if err := json.NewDecoder(response.Body).Decode(&invoices); err != nil {
		return nil, 0, 0, fmt.Errorf("error decoding response: %w", err)
	}

	var resultInvoices []model.InvoiceListResponse
	invoiceMap := make(map[int64]model.InvoiceListResponse)
	var totalInvoicesCMS int
	var totalAmountCMS float64

	for _, invoice := range invoices.Data {
		roundPrice := math.Round(invoice.InUsd*100) / 100
		if existingInvoice, exists := invoiceMap[invoice.ID]; exists {
			// Si la factura ya existe, sumar InUsd y combinar los items
			existingInvoice.InUsd += roundPrice
			totalAmountCMS += roundPrice

			if invoice.AlegraTransactionList.InvoiceRelationList.InvoiceItems != nil {
				for _, item4 := range invoice.AlegraTransactionList.InvoiceRelationList.InvoiceItems {
					invoiceItem := model.InvoiceItemListResponse{
						OriginalPrice: item4.OriginalPrice,
					}
					existingInvoice.AlegraTransactionListResponse.InvoiceRelationListResponse.InvoiceItems = append(existingInvoice.AlegraTransactionListResponse.InvoiceRelationListResponse.InvoiceItems, invoiceItem)
				}
			}
			invoiceMap[invoice.ID] = existingInvoice
			continue
		}
		// Si es una nueva factura, crearla
		var invoiceItems []model.InvoiceItemListResponse

		if invoice.AlegraTransactionList.InvoiceRelationList.InvoiceItems != nil {
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
			ID:            invoice.ID,
			UserID:        int64(invoice.UserID),
			Email:         invoice.Email,
			FirstName:     invoice.FirstName,
			LastName:      invoice.LastName,
			InUsd:         roundPrice,
			ExchangeRate:  invoice.ExchangeRate,
			OriginalPrice: invoice.OriginalPrice,

			AlegraTransactionListResponse: alegraTransaction,
		}

		invoiceMap[invoice.ID] = newInvoice
		totalAmountCMS += roundPrice

	}

	for _, invoice := range invoiceMap {
		resultInvoices = append(resultInvoices, invoice)
		totalInvoicesCMS++
	}

	jsonData, err := json.Marshal(resultInvoices)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("error marshalling result invoices: %w", err)
	}

	totalAmountCMS = math.Round((totalAmountCMS)*100) / 100

	return jsonData, totalInvoicesCMS, totalAmountCMS, nil
}
