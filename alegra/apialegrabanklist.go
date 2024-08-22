package alegra

import (
	"cmsalegra/configuration"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func IdsBankAccounts(config configuration.Configuration) ([]int, error) {

	authAlegraToken, err := basicAuthAlegraAPI(config.ALEGRAApi.AlegraEmail, config.ALEGRAApi.AlegraToken)
	if err != nil {
		return nil, err
	}

	baseURL := fmt.Sprintf("%s/api/v1/bank-accounts", config.ALEGRAApi.UrlApiAlegra)

	params := url.Values{}
	params.Add("metadata", "true")

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error al crear la solicitud HTTP: %v", err)
	}

	req.Header.Set("Authorization", authAlegraToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error al realizar la solicitud HTTP: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer el cuerpo de la respuesta: %v", err)
	}

	var response ResponseBankAccounts

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error al decodificar la respuesta JSON: %v", err)
	}

	var ids []int

	// Iterar sobre los elementos en Data y convertir los IDs a int
	for _, item := range response.Data {
		idInt, err := strconv.Atoi(item.ID)
		if err != nil {
			return nil, fmt.Errorf("error al convertir ID a entero: %v", err)
		}
		ids = append(ids, idInt)
	}

	return ids, nil
}

type Metadata struct {
	Total int `json:"total"`
}

type BankAccount struct {
	ID string `json:"id"`
}

type ResponseBankAccounts struct {
	Metadata Metadata      `json:"metadata"`
	Data     []BankAccount `json:"data"`
}
