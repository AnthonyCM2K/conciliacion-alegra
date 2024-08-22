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
}

type AlegraTransactionListResponse struct {
	InvoiceRelationListResponse InvoiceRelationListResponse `json:"invoice_relation"`
	AlegraPaymentID             string                      `json:"alegra_payment_id"`
}

type InvoiceRelationListResponse struct {
	InvoiceItems []InvoiceItemListResponse `json:"invoice_items"`
}
type InvoiceItemListResponse struct {
	OriginalPrice float64 `json:"original_price"`
}
