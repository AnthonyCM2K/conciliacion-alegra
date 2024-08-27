package main

import (
	"bufio"
	"cmsalegra/conciliacion"
	"cmsalegra/configuration"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

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

	var date string

	if *datePtr != "" {
		date = *datePtr
	} else {
		// Solicitar input de la fecha si no se proporcionó el flag
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Por favor, ingrese la fecha en formato YYYY-MM-DD: ")
		inputDate, _ := reader.ReadString('\n')
		date = strings.TrimSpace(inputDate)
	}

	// Validar la fecha
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		fmt.Println(`Error: La fecha proporcionada no tiene el formato correcto. Debe ser YYYY-MM-DD ejemplo: 2024-08-19`)
		return
	}

	conciliacion.Conciliation(date, config)
}
