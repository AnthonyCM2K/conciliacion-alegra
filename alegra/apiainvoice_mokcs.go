package alegra

import (
	"cmsalegra/configuration"
	"cmsalegra/model"
	"time"
)

var mockConfig = configuration.Configuration{
	ALEGRAApi: configuration.AlegraApi{
		AlegraEmail:  "test@example.com",
		AlegraToken:  "test_token",
		UrlApiAlegra: "http://localhost:8080",
	},
}

var mockPayments = []model.Data{
	{
		ID:               "6440",
		Date:             model.CustomDate(time.Date(2024, 9, 10, 0, 0, 0, 0, time.UTC)),
		Number:           "4373",
		Amount:           172.01,
		Anotation:        "AMEX",
		Type:             "in",
		Status:           "open",
		DecimalPrecision: "2",
		CalculationScale: "2",
		BankAccount: model.BankAccount{
			ID:   "4",
			Name: "Banco Interbank Dolares",
			Type: "bank",
			Currency: model.Currency{
				Symbol: "$",
			},
		},
		Client: model.Client{},
		Currency: model.Currency{
			Code:         "USD",
			Symbol:       "$",
			ExchangeRate: 3.9,
		},
		Categories: []model.Category{
			{ID: "5093"}, // Categoría no válida
		},
		NumberTemplate: model.NumberTemplate{
			ID:              "5",
			Prefix:          "",
			Number:          "4373",
			FullNumber:      "4373",
			FormattedNumber: "4373",
		},
	},
	{
		ID:               "6434",
		Date:             model.CustomDate(time.Date(2024, 9, 10, 0, 0, 0, 0, time.UTC)),
		Number:           "4372",
		Amount:           1.05,
		Type:             "in",
		PaymentMethod:    "transfer",
		Status:           "open",
		DecimalPrecision: "2",
		CalculationScale: "2",
		BankAccount: model.BankAccount{
			ID:   "9",
			Name: "Pasarela Ebanx",
			Type: "bank",
			Currency: model.Currency{
				Symbol: "$",
			},
		},
		Client: model.Client{
			ID:             "169",
			Name:           "Varios",
			Identification: "Varios",
		},
		Currency: model.Currency{
			Code:         "USD",
			Symbol:       "$",
			ExchangeRate: 3.9,
		},
		Categories: []model.Category{
			{ID: "5283"}, // Categoría válida
		},
		NumberTemplate: model.NumberTemplate{
			ID:              "5",
			Prefix:          "",
			Number:          "4372",
			FullNumber:      "4372",
			FormattedNumber: "4372",
		},
	},
	{
		ID:               "6421",
		Date:             model.CustomDate(time.Date(2024, 9, 10, 0, 0, 0, 0, time.UTC)),
		Number:           "4365",
		Amount:           300,
		Anotation:        "Valor pagado: 117.000000 PEN (1 USD = 3.900000 PEN)",
		Type:             "in",
		PaymentMethod:    "credit-card",
		Status:           "open",
		DecimalPrecision: "2",
		CalculationScale: "2",
		BankAccount: model.BankAccount{
			ID:       "10",
			Name:     "Pasarela Niubiz soles",
			Type:     "bank",
			Currency: model.Currency{},
		},
		Client: model.Client{
			ID:             "247",
			Name:           "Clientes no domiciliados",
			Identification: "V0001",
		},
		Currency: model.Currency{
			Code:         "USD",
			Symbol:       "$",
			ExchangeRate: 3.9,
		},
		Categories: []model.Category{
			{ID: "5324"}, // Categoría válida
		},
		NumberTemplate: model.NumberTemplate{
			ID:              "5",
			Prefix:          "",
			Number:          "4365",
			FullNumber:      "4365",
			FormattedNumber: "4365",
		},
	},
}
