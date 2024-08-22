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
func requestCMS2(fecha string, config configuration.Configuration) (*http.Response, error) {

	client := &http.Client{}

	//Peticion de token para CMS
	tokenCMS, err := ApiCMSLogin(config)
	if err != nil {
		log.Fatal(err)
	}

	// Base URL
	baseURL := config.CMSApi.UrlApiCmsConsulta2

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

func QueryApiByteCMSReports(fecha string, config configuration.Configuration) ([]byte, error) {

	response, err := requestCMS2(fecha, config)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var invoices model.InvoiceCMSList
	if err := json.NewDecoder(response.Body).Decode(&invoices); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	var resultInvoices []model.InvoiceListResponse

	for _, invoice := range invoices.Data {

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
			ID:        invoice.ID,
			UserID:    int64(invoice.UserID),
			Email:     invoice.Email,
			FirstName: invoice.FirstName,
			LastName:  invoice.LastName,

			AlegraTransactionListResponse: alegraTransaction,
		}

		resultInvoices = append(resultInvoices, newInvoice)
	}

	jsonData, err := json.Marshal(resultInvoices)
	if err != nil {
		return nil, fmt.Errorf("error marshalling result invoices: %w", err)
	}

	return jsonData, nil

}
