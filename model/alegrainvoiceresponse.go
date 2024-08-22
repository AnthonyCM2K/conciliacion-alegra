package model

type FacturasAlegraResponse struct {
	Metadata MetadataResponse        `json:"metadata,omitempty"`
	Data     []InvoiceAlegraResponse `json:"data,omitempty"`
}

type InvoiceAlegraResponse struct {
	ID       int64              `json:"id"`
	Amount   float64            `json:"amount"`
	Currency []CurrencyResponse `json:"currency,omitempty"`
}

type ClientResponse struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Identification string `json:"identification,omitempty"`
}

type CurrencyResponse struct {
	Code         string  `json:"code,omitempty"`
	Symbol       string  `json:"symbol,omitempty"`
	ExchangeRate float64 `json:"exchangeRate,omitempty"`
}

type MetadataResponse struct {
	Total int64 `json:"total,omitempty"`
}

// type InvoiceAlegraResponse struct {
// 	ID               string           `json:"id,omitempty"`
// 	Date             CustomDate       `json:"date,omitempty"`
// 	Amount           string           `json:"amount,omitempty"`
// 	Anotation        interface{}      `json:"anotation"`
// 	Number           string           `json:"number,omitempty"`
// 	Status           string           `json:"status,omitempty"`
// 	Type             string           `json:"type,omitempty"`
// 	Client           ClientResponse   `json:"client,omitempty"`
// 	Currency         CurrencyResponse `json:"currency,omitempty"`
// 	Subtype          string           `json:"subtype,omitempty"`
// 	Conciliation     interface{}      `json:"conciliation"`
// 	Conciliated      bool             `json:"conciliated,omitempty"`
// 	DecimalPrecision string           `json:"decimalPrecision,omitempty"`
// 	Deletable        bool             `json:"deletable,omitempty"`
// 	Voidable         bool             `json:"voidable,omitempty"`
// 	Openable         bool             `json:"openable,omitempty"`
// 	MovementType     string           `json:"movementType,omitempty"`
// 	IsMinorExpense   bool             `json:"isMinorExpense,omitempty"`
// 	Editable         bool             `json:"editable,omitempty"`
// 	Associations     string           `json:"associations,omitempty"`
// }
