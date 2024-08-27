package alegra

import (
	"cmsalegra/configuration"
	"cmsalegra/model"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"

	"golang.org/x/exp/slices"
)

func basicAuthAlegraAPI(username, password string) (string, error) {
	auth := username + ":" + password
	basicAuth := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(auth)))

	return basicAuth, nil
}

func RequestAlegraPayments(fecha string, config configuration.Configuration) ([]model.Data, error) {
	var allPayments []model.Data
	start := 0
	limit := 30

	authAlegraToken, err := basicAuthAlegraAPI(config.ALEGRAApi.AlegraEmail, config.ALEGRAApi.AlegraToken)
	if err != nil {
		return nil, err
	}

	// Definir los parámetros de consulta estáticos usando `url.Values`
	params := url.Values{}
	params.Add("order_direction", "DESC")
	params.Add("order_field", "date")
	params.Add("type", "in")
	params.Add("date", fecha)

	// URL base de la API
	baseURL := config.ALEGRAApi.UrlApiAlegra + "/api/v1/payments"

	for {
		// Agregar parámetros de paginación
		params.Set("start", fmt.Sprintf("%d", start))
		params.Set("limit", fmt.Sprintf("%d", limit))

		// Construir la URL completa con los parámetros
		fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

		// Crear una nueva solicitud HTTP
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return nil, err
		}

		// Agregar el token en el header Authorization
		req.Header.Set("Authorization", authAlegraToken)

		// Crear un cliente HTTP y realizar la solicitud
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		// Verificar el estado de la respuesta
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("request failed with status: %s", resp.Status)
		}

		// Decodificar la respuesta JSON directamente a un array de pagos
		var payments []model.Data
		err = json.NewDecoder(resp.Body).Decode(&payments)
		if err != nil {
			return nil, err
		}

		// Agregar los pagos obtenidos a la lista total
		allPayments = append(allPayments, payments...)

		// Verificar si se recibieron menos registros que el límite, para detener la paginación
		if len(payments) < limit {
			break
		}

		start += limit
	}

	return allPayments, nil
}

func QueryApiByteAlegra(fecha string, config configuration.Configuration) ([]byte, int, float64, error) {
	// Obtener los pagos desde la API de Alegra
	payments, err := RequestAlegraPayments(fecha, config)
	if err != nil {
		return nil, 0, 0, err
	}

	var resultInvoices []model.InvoiceAlegraResponse
	var totalInvoicesAlegra int
	var totalAmountAlegra float64

	// Procesar y transformar los datos obtenidos
	for _, payment := range payments {

		//Filtrar por categoria, es transferencia
		isNotAnInvoice := (len(payment.Categories) >= 1 && !slices.Contains(model.Categories, payment.Categories[0].ID)) || payment.Type == "out"
		if isNotAnInvoice {
			continue
		}

		// Sumar al total de facturas
		totalInvoicesAlegra++

		// Sumar al total de Amount
		totalAmountAlegra += payment.Amount

		id, err := strconv.ParseInt(payment.ID, 10, 64)
		if err != nil {
			log.Printf("Error parsing invoice ID: %v", err)
			continue
		}
		// Convierte el Amount a float64 para DEV
		// amount, err := strconv.ParseFloat(payment.Amount, 64)
		// if err != nil {
		// 	log.Printf("Error parsing decimal precision: %v", err)
		// 	continue
		// }

		// Crear la respuesta transformada
		newInvoice := model.InvoiceAlegraResponse{
			ID:     id,
			Amount: payment.Amount,
			Currency: []model.CurrencyResponse{
				{
					Code:         payment.Currency.Code,
					ExchangeRate: payment.Currency.ExchangeRate,
				},
			},
			BankAccount: model.BankAccountResponse{
				Name: payment.BankAccount.Name,
			},
		}
		resultInvoices = append(resultInvoices, newInvoice)

	}

	jsonData, err := json.Marshal(resultInvoices)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("error marshalling result invoices: %w", err)
	}

	totalAmountAlegra = math.Round((totalAmountAlegra)*100) / 100

	return jsonData, totalInvoicesAlegra, totalAmountAlegra, nil
}
