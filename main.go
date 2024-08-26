package main

import (
	"cmsalegra/conciliacion"
	"cmsalegra/configuration"
	"flag"
	"fmt"
	"time"
)

// "cmsalegra/conciliacion"
// "cmsalegra/configuration"
// "flag"
// "fmt"
// "time"

func main() {
	defaultFilePath := "configuration.json"

	// Definir el flag "filePath" opcional
	filePathPtr := flag.String("filePath", defaultFilePath, "Ruta al archivo de configuration.json")

	// Definir el flag "date" para recibir una fecha como string
	datePtr := flag.String("date", "", "Fecha en formato YYYY-MM-DD")

	flag.Parse()

	// Usar el valor del flag "filePath"
	filePath := *filePathPtr

	config := configuration.NewConfiguration(filePath)

	if *datePtr != "" {
		_, err := time.Parse("2006-01-02", *datePtr)
		if err != nil {
			fmt.Println(`Error: La fecha proporcionada en el el flag -date no tiene el formato correcto. Debe ser YYYY-MM-DD ejemplo: -date="2024-08-19"`)
			return
		}

		conciliacion.Conciliation(*datePtr, config)
	} else {
		fmt.Println(`Debe proporcionar el flag -date con una fecha en formato YYYY-MM-DD ejemplo: -date="2024-08-19"`)
	}
}

// func main() {
// 	fecha := "2024-08-22"

// 	config := configuration.NewConfiguration("configuration.json")

// 	payments, err := alegra.QueryApiByteAlegra(fecha, config)
// 	if err != nil {
// 		fmt.Println("Error al obtener los pagos:", err)
// 		return
// 	}
// 	// Aquí puedes usar la lista completa de pagos (payments) como desees
// 	fmt.Println("Total de pagos obtenidos:", string(payments))
// 	// Exportar o procesar la lista completa de pagos
// }
