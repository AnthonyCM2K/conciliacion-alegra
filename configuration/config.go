package configuration

import (
	"encoding/json"
	"log"
	"os"
)

// Configuration model
type Configuration struct {
	CMSApi    CmsApi    `json:"cmsapi"`
	ALEGRAApi AlegraApi `json:"alegraapi"`
}
type AlegraApi struct {
	UrlApiAlegra string `json:"urlapialegra"`
	AlegraEmail  string `json:"alegraemail"`
	AlegraToken  string `json:"alegratoken"`
}
type CmsApi struct {
	UrlApiCmsConsulta      string `json:"urlapicmsconsulta"`
	UrlApiCmsLogin         string `json:"urlapicmslogin"`
	UrlApiCmsLoginEmail    string `json:"urlapicmslogin_email"`
	UrlApiCmsLoginPassword string `json:"urlapicmslogin_password"`
}

func NewConfiguration(path string) Configuration {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	conf := Configuration{}
	if err := json.Unmarshal(file, &conf); err != nil {
		log.Fatal(err)
	}

	return conf
}

// func RoundToTwoDecimalPlaces(value float64) float64 {
// 	return math.Round(value*100) / 100
// }
