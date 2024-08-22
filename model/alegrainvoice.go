package model

import "time"

type FacturasAlegra3 struct {
	Metadata Metadata3 `json:"metadata,omitempty"`
	Data     []Data3   `json:"data,omitempty"`
}

type Data3 struct {
	ID               string      `json:"id,omitempty"`
	Date             CustomDate  `json:"date,omitempty"`
	Amount           string      `json:"amount,omitempty"`
	Anotation        interface{} `json:"anotation"`
	Number           string      `json:"number,omitempty"`
	Status           string      `json:"status,omitempty"`
	Type             string      `json:"type,omitempty"`
	Client           Client3     `json:"client,omitempty"`
	Currency         Currency3   `json:"currency,omitempty"`
	Subtype          string      `json:"subtype,omitempty"`
	Conciliation     interface{} `json:"conciliation"`
	Conciliated      bool        `json:"conciliated,omitempty"`
	DecimalPrecision string      `json:"decimalPrecision,omitempty"`
	Deletable        bool        `json:"deletable,omitempty"`
	Voidable         bool        `json:"voidable,omitempty"`
	Openable         bool        `json:"openable,omitempty"`
	MovementType     string      `json:"movementType,omitempty"`
	IsMinorExpense   bool        `json:"isMinorExpense,omitempty"`
	Editable         bool        `json:"editable,omitempty"`
	Associations     string      `json:"associations,omitempty"`
}

type Client3 struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Identification string `json:"identification,omitempty"`
}

type Currency3 struct {
	Code         string  `json:"code,omitempty"`
	Symbol       string  `json:"symbol,omitempty"`
	ExchangeRate float64 `json:"exchangeRate,omitempty"`
}

type Metadata3 struct {
	Total int64 `json:"total,omitempty"`
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
