package model

import "time"

type InvoiceCMSList struct {
	Data     []InvoiceList `json:"data"`
	Messages []Message     `json:"messages"`
}

type Message struct {
	Code string `json:"code"`
}

type InvoiceList struct {
	ID                    int64                  `json:"id"`
	InvoiceDate           time.Time              `json:"invoice_date"`
	UserID                int64                  `json:"user_id"`
	Email                 string                 `json:"email"`
	FirstName             string                 `json:"first_name"`
	LastName              string                 `json:"last_name"`
	CashBoxID             int64                  `json:"cash_box_id"`
	Responsable           string                 `json:"responsable"`
	Account               string                 `json:"account"`
	State                 string                 `json:"state"`
	Support               string                 `json:"support"`
	AlegraTransactionList *AlegraTransactionList `json:"alegra_transaction,omitempty"`
	IsSubscription        bool                   `json:"is_subscription"`
	IsRenewal             bool                   `json:"is_renewal"`
	Months                int64                  `json:"months"`
	BeginsAt              time.Time              `json:"begins_at"`
	EndsAt                time.Time              `json:"ends_at"`
	CourseID              int64                  `json:"course_id"`
	Course                string                 `json:"course"`
	CurrencyID            int64                  `json:"currency_id"`
	Currency              string                 `json:"currency"`
	ExchangeRate          float64                `json:"exchange_rate"`
	OriginalPrice         float64                `json:"original_price"`
	DiscountPrice         float64                `json:"discount_price"`
	InUsd                 float64                `json:"in_usd"`
	CreatedAt             time.Time              `json:"created_at"`
	Country               string                 `json:"country"`
	AddressCountry        string                 `json:"address_country"`
	PaymentMethod         string                 `json:"payment_method"`
	PaymentMethodType     string                 `json:"payment_method_type"`
}

type AlegraTransactionList struct {
	Status              string              `json:"status"`
	CurrencyList        CurrencyList        `json:"currency"`
	Anotation           string              `json:"anotation"`
	Categories          []CategoryList      `json:"categories"`
	AlegraDataList      AlegraDataList      `json:"alegra_data"`
	InvoiceRelationList InvoiceRelationList `json:"invoice_relation"`
	AlegraPaymentID     string              `json:"alegra_payment_id"`
}

type AlegraDataList struct {
	ID            int64  `json:"id"`
	BankAccount   string `json:"bank_account"`
	PaymentMethod string `json:"payment_method"`
}

type CategoryList struct {
	ID           int64   `json:"id"`
	Tax          Tax     `json:"tax"`
	Price        float64 `json:"price"`
	Quantity     float64 `json:"quantity"`
	Observations string  `json:"observations"`
}

type CurrencyList struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Description  string    `json:"description"`
	ExchangeRate float64   `json:"exchange_rate"`
}

type InvoiceRelationList struct {
	CurrencyList  CurrencyList      `json:"currency"`
	InvoiceItems  []InvoiceItemList `json:"invoice_items"`
	InvoiceHeader InvoiceHeaderList `json:"invoice_header"`
}

type InvoiceHeaderList struct {
	ID                  int64       `json:"id"`
	State               string      `json:"state"`
	DetailList          DetailList  `json:"detail"`
	Support             string      `json:"support"`
	UserID              int64       `json:"user_id"`
	GroupID             int64       `json:"group_id"`
	OrderID             int64       `json:"order_id"`
	Sequence            int64       `json:"sequence"`
	AlegraID            string      `json:"alegra_id"`
	CashboxID           int64       `json:"cashbox_id"`
	CreatedAt           time.Time   `json:"created_at"`
	UpdatedAt           time.Time   `json:"updated_at"`
	ApprovedBy          int64       `json:"approved_by"`
	InvoiceDate         time.Time   `json:"invoice_date"`
	Observations        string      `json:"observations"`
	ProratedDays        int64       `json:"prorated_days"`
	PaymentMethod       string      `json:"payment_method"`
	AlegraTransaction   interface{} `json:"alegra_transaction"`
	SubscriptionOrderID int64       `json:"subscription_order_id"`
}

type DetailList struct {
	TotalsList                     TotalsList                     `json:"totals"`
	CreatedAtUser                  time.Time                      `json:"created_at_user"`
	FirstSaleDate                  time.Time                      `json:"first_sale_date"`
	GroupCreatedAt                 time.Time                      `json:"group_created_at"`
	IsFirstPurchase                bool                           `json:"is_first_purchase"`
	DetailsBeforeFirstPurchaseList DetailsBeforeFirstPurchaseList `json:"details_before_first_purchase"`
	DaysElapsedPreviousPurchase    int64                          `json:"days_elapsed_previous_purchase"`
}

type DetailsBeforeFirstPurchaseList struct {
	DaysElapsed   int64 `json:"days_elapsed"`
	SubjectsViews int64 `json:"subjects_views"`
}

type TotalsList struct {
	Money   float64 `json:"money"`
	Tickets int64   `json:"tickets"`
}

type InvoiceItemList struct {
	ID                  int64     `json:"id"`
	Quantity            float64   `json:"quantity"`
	CourseID            int64     `json:"course_id"`
	BasePrice           float64   `json:"base_price"`
	CreatedAt           time.Time `json:"created_at"`
	InvoiceID           int64     `json:"invoice_id"`
	UnitPrice           float64   `json:"unit_price"`
	UpdatedAt           time.Time `json:"updated_at"`
	CurrencyID          int64     `json:"currency_id"`
	ExchangeRate        float64   `json:"exchange_rate"`
	DiscountPrice       float64   `json:"discount_price"`
	OriginalPrice       float64   `json:"original_price"`
	IsSubscription      bool      `json:"is_subscription"`
	IsAttendingCourse   bool      `json:"is_attending_course"`
	ProfessorPercentage float64   `json:"professor_percentage"`
}

type Tax struct {
	ID int64 `json:"id"`
}
