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
	UrlApiCmsConsulta2     string `json:"urlapicmsconsulta"`
	UrlApiCmsLogin         string `json:"urlapicmslogin"`
	UrlApiCmsLoginEmail    string `json:"urlapicmslogin_email"`
	UrlApiCmsLoginPassword string `json:"urlapicmslogin_password"`
}

// Config almacena la configuración cargada.
// var Config Configuration

// func init() {
// 	filePath := "/home/anthonylinux/Documentos/0_EDTEAM/cmsalegra/configuration.json"
// 	file, err := os.ReadFile(filePath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if err := json.Unmarshal(file, &Config); err != nil {
// 		log.Fatal(err)
// 	}
// }

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
