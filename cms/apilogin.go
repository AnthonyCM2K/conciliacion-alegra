package cms

import (
	"bytes"
	"cmsalegra/configuration"
	"encoding/json"
	"fmt"

	"io"
	"net/http"
)

type CMSClient interface {
	ApiCMSLogin(config configuration.Configuration) (string, error)
}

type DefaultCMSClient struct{}

func (d *DefaultCMSClient) ApiCMSLogin(config configuration.Configuration) (string, error) {

	url := config.CMSApi.UrlApiCmsLogin
	credentials := map[string]string{
		"email":    config.CMSApi.UrlApiCmsLoginEmail,
		"password": config.CMSApi.UrlApiCmsLoginPassword,
	}

	jsonData, err := json.Marshal(credentials)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)

	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(err)
	}

	data, ok := result["data"].(map[string]interface{})
	if !ok {
		fmt.Println("Estructura de respuesta inesperada")
	}

	token, ok := data["token"].(string)
	if !ok {
		fmt.Println("Token no encontrado en la respuesta")
	}

	return token, nil

}
