package alegra

import (
	"cmsalegra/configuration"
	"cmsalegra/model"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func basicAuthAlegraAPI(username, password string) (string, error) {
	auth := username + ":" + password
	basicAuth := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(auth)))

	return basicAuth, nil
}

func requestAlegraPayments(fecha string, id int, config configuration.Configuration) (*model.FacturasAlegra3, error) {
	authAlegraToken, err := basicAuthAlegraAPI(config.ALEGRAApi.AlegraEmail, config.ALEGRAApi.AlegraToken)
	if err != nil {
		return nil, err
	}

	baseURL := fmt.Sprintf("%s/api/v1/bank-accounts/%%d/payments", config.ALEGRAApi.UrlApiAlegra)

	fullURL := fmt.Sprintf(baseURL, id)

	// Convertir la fecha al formato requerido
	fechaConvertida := convertirFechaRecibida(fecha)

	// Parámetros de consulta
	params := url.Values{}
	//params.Add("limit", `3`)
	params.Add("order_direction", `DES`)
	//params.Add("fields", `conciliation,editable,deletable,associations`)
	params.Add("date", fechaConvertida) //`2024%2F08%2F16`

	fullURLWithParams := fmt.Sprintf("%s?%s", fullURL, params.Encode())

	req, err := http.NewRequest("GET", fullURLWithParams, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", authAlegraToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Lee el cuerpo de la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Decodifica la respuesta JSON
	var payment model.FacturasAlegra3
	if err := json.Unmarshal(body, &payment); err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return &payment, nil
}

func QueryApiByteAlegraBankAll(fecha string, config configuration.Configuration) ([]byte, error) {
	// Llama a la función fetchBankAccountIDs para obtener los IDs
	ids, err := IdsBankAccounts(config)
	if err != nil {
		return nil, fmt.Errorf("error al obtener los IDs: %w", err)
	}

	var allPayments []model.InvoiceAlegraResponse

	for _, id := range ids {
		invoices, err := requestAlegraPayments(fecha, id, config)
		if err != nil {
			log.Printf("Error fetching payments for ID %d: %v", id, err)
			continue
		}

		// Procesar y transformar los datos
		for _, invoice := range invoices.Data {
			// Convierte el ID a int64
			id, err := strconv.ParseInt(invoice.ID, 10, 64)
			if err != nil {
				log.Printf("Error parsing invoice ID: %v", err)
				continue
			}
			// Convierte el Amount a float64
			amount, err := strconv.ParseFloat(invoice.Amount, 64)
			if err != nil {
				log.Printf("Error parsing decimal precision: %v", err)
				continue
			}

			// Crea la respuesta transformada
			newInvoice := model.InvoiceAlegraResponse{
				ID:     id,
				Amount: amount,
				Currency: []model.CurrencyResponse{
					{
						Code:         invoice.Currency.Code,
						ExchangeRate: invoice.Currency.ExchangeRate,
					},
				},
			}
			allPayments = append(allPayments, newInvoice)
		}
	}

	jsonData, err := json.Marshal(allPayments)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payments to JSON: %w", err)
	}

	return jsonData, nil
}

// convertirFechaRecibida reemplaza los guiones en la fecha con "%2F".
func convertirFechaRecibida(fecha string) string {
	return strings.ReplaceAll(fecha, "-", "%2F")
}
