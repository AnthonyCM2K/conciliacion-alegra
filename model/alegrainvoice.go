package model

import "time"

var Categories = []string{"5283", "5284", "5285", "5286", "5324", "5325", "5356", "5358", "5363"}

type FacturasAlegra struct {
	Metadata Metadata `json:"metadata,omitempty"`
	Data     []Data   `json:"data,omitempty"`
}

type Data struct {
	ID               string         `json:"id,omitempty"`
	Date             CustomDate     `json:"date,omitempty"`
	Number           string         `json:"number,omitempty"`
	Amount           float64        `json:"amount,omitempty"`
	Observations     string         `json:"observations,omitempty"`
	Anotation        string         `json:"anotation,omitempty"`
	Type             string         `json:"type,omitempty"`
	PaymentMethod    string         `json:"paymentMethod,omitempty"`
	Status           string         `json:"status,omitempty"`
	DecimalPrecision string         `json:"decimalPrecision,omitempty"`
	CalculationScale string         `json:"calculationScale,omitempty"`
	BankAccount      BankAccount    `json:"bankAccount,omitempty"`
	Client           Client         `json:"client,omitempty"`
	Currency         Currency       `json:"currency,omitempty"`
	Categories       []Category     `json:"categories,omitempty"`
	CostCenter       string         `json:"costCenter,omitempty"`
	NumberTemplate   NumberTemplate `json:"numberTemplate,omitempty"`
}
type BankAccount struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Type     string   `json:"type"`
	Currency Currency `json:"currency"`
}

type Category struct {
	ID string `json:"id,omitempty"`
	// Name         string   `json:"name"`
	// Price        float64  `json:"price"`
	// Quantity     int      `json:"quantity"`
	// Observations string   `json:"observations"`
	// Total        float64  `json:"total"`
	// Behavior     string   `json:"behavior"`
	// Tax          []string `json:"tax"`
}

type Client struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Phone          string `json:"phone,omitempty"`
	Identification string `json:"identification,omitempty"`
}

type Currency struct {
	Code         string  `json:"code,omitempty"`
	Symbol       string  `json:"symbol,omitempty"`
	ExchangeRate float64 `json:"exchangeRate,omitempty"`
}

type Metadata struct {
	Total int64 `json:"total,omitempty"`
}

type NumberTemplate struct {
	ID              string `json:"id"`
	Prefix          string `json:"prefix"`
	Number          string `json:"number"`
	FullNumber      string `json:"fullNumber"`
	FormattedNumber string `json:"formattedNumber"`
}

// CustomDate representa una fecha personalizada
type CustomDate time.Time

// UnmarshalJSON implementa la deserialización personalizada para CustomDate
func (cd *CustomDate) UnmarshalJSON(b []byte) error {
	const layout = "2006-01-02"
	s := string(b)
	t, err := time.Parse(`"`+layout+`"`, s)
	if err != nil {
		return err
	}
	*cd = CustomDate(t)
	return nil
}
