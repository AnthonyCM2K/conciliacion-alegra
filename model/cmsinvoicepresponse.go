package model

type InvoiceCMSListResponse struct {
	Data     []InvoiceListResponse `json:"data"`
	Messages []Message             `json:"messages"`
}

type InvoiceListResponse struct {
	ID                            int64                         `json:"id"`
	UserID                        int64                         `json:"user_id"`
	Email                         string                        `json:"email"`
	FirstName                     string                        `json:"first_name"`
	LastName                      string                        `json:"last_name"`
	AlegraTransactionListResponse AlegraTransactionListResponse `json:"alegra_transaction,omitempty"`
	InUsd                         float64                       `json:"in_usd"`
	ExchangeRate                  float64                       `json:"exchange_rate"`
	Currency                      string                        `json:"currency"`
	PaymentMethod                 string                        `json:"payment_method"`
	OriginalPrice                 float64                       `json:"original_price"`
}

type AlegraTransactionListResponse struct {
	InvoiceRelationListResponse InvoiceRelationListResponse `json:"invoice_relation"`
	AlegraPaymentID             string                      `json:"alegra_payment_id"`
	AlegraDataList              AlegraDataListResponse      `json:"alegra_data"`
}

type InvoiceRelationListResponse struct {
	InvoiceItems []InvoiceItemListResponse `json:"invoice_items"`
}
type InvoiceItemListResponse struct {
	OriginalPrice float64 `json:"original_price"`
}

type AlegraDataListResponse struct {
	//ID            int64  `json:"id"`
	BankAccount string `json:"bank_account"`
	//PaymentMethod string `json:"payment_method"`
}
