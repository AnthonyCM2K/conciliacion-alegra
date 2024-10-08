package model

type FacturasAlegraResponse struct {
	Metadata MetadataResponse        `json:"metadata,omitempty"`
	Data     []InvoiceAlegraResponse `json:"data,omitempty"`
}

type InvoiceAlegraResponse struct {
	ID       int64              `json:"id"`
	Amount   float64            `json:"amount"`
	Currency []CurrencyResponse `json:"currency,omitempty"`
	//PaymentMethod string             `json:"paymentMethod"`
	BankAccount BankAccountResponse `json:"bankAccount"`
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

type BankAccountResponse struct {
	//ID       string    `json:"id"`
	Name string `json:"name"`
	//Type     string    `json:"type"`
	//Currency Currency3 `json:"currency"`
}
