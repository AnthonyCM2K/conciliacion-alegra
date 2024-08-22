package model

type InvoiceCMSListResponse struct {
	Data     []InvoiceListResponse `json:"data"`
	Messages []Message             `json:"messages"`
}

type InvoiceListResponse struct {
	ID int64 `json:"id"`
	//InvoiceDate           time.Time             `json:"invoice_date"`
	UserID    int64  `json:"user_id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	//CashBoxID int64  `json:"cash_box_id"`
	//Responsable           string                `json:"responsable"`
	//Account               string                `json:"account"`
	//State                 string                `json:"state"`
	//Support               string                `json:"support"`
	AlegraTransactionListResponse AlegraTransactionListResponse `json:"alegra_transaction,omitempty"`
	//IsSubscription        bool                  `json:"is_subscription"`
	//IsRenewal             bool                  `json:"is_renewal"`
	//Months                int64                 `json:"months"`
	//BeginsAt              time.Time             `json:"begins_at"`
	// EndsAt                time.Time             `json:"ends_at"`
	// CourseID              int64                 `json:"course_id"`
	// Course                string                `json:"course"`
	// CurrencyID            int64                 `json:"currency_id"`
	// Currency              string                `json:"currency"`
	// ExchangeRate          int64                 `json:"exchange_rate"`
	// OriginalPrice         int64                 `json:"original_price"`
	// DiscountPrice         int64                 `json:"discount_price"`
	InUsd float64 `json:"in_usd"`
	// CreatedAt             time.Time             `json:"created_at"`
	// Country               string                `json:"country"`
	// AddressCountry        string                `json:"address_country"`
	// PaymentMethod         string                `json:"payment_method"`
	// PaymentMethodType     string                `json:"payment_method_type"`
}

type AlegraTransactionListResponse struct {
	// Status              string              `json:"status"`
	// CurrencyList        CurrencyList        `json:"currency"`
	// Anotation           string              `json:"anotation"`
	// Categories          []CategoryList      `json:"categories"`
	// AlegraDataList      AlegraDataList      `json:"alegra_data"`
	InvoiceRelationListResponse InvoiceRelationListResponse `json:"invoice_relation"`
	AlegraPaymentID             string                      `json:"alegra_payment_id"`
}

type InvoiceRelationListResponse struct {
	// CurrencyList  CurrencyList      `json:"currency"`
	InvoiceItems []InvoiceItemListResponse `json:"invoice_items"`
	// InvoiceHeaderList InvoiceHeaderList     `json:"invoice_header"`
}

// type AlegraDataList struct {
// 	ID            int64  `json:"id"`
// 	BankAccount   string `json:"bank_account"`
// 	PaymentMethod string `json:"payment_method"`
// }

// type CategoryList struct {
// 	ID           int64  `json:"id"`
// 	Tax          Tax    `json:"tax"`
// 	Price        int64  `json:"price"`
// 	Quantity     int64  `json:"quantity"`
// 	Observations string `json:"observations"`
// }

// type CurrencyList struct {
// 	ID           int64     `json:"id"`
// 	Name         string    `json:"name"`
// 	CreatedAt    time.Time `json:"created_at"`
// 	UpdatedAt    time.Time `json:"updated_at"`
// 	Description  string    `json:"description"`
// 	ExchangeRate float64   `json:"exchange_rate"`
// }

// type InvoiceHeaderList struct {
// 	ID                  int64              `json:"id"`
// 	State               string             `json:"state"`
// 	DetailListResponse  DetailListResponse `json:"detail"`
// 	Support             string             `json:"support"`
// 	UserID              int64              `json:"user_id"`
// 	GroupID             int64              `json:"group_id"`
// 	OrderID             int64              `json:"order_id"`
// 	Sequence            int64              `json:"sequence"`
// 	AlegraID            string             `json:"alegra_id"`
// 	CashboxID           int64              `json:"cashbox_id"`
// 	CreatedAt           time.Time          `json:"created_at"`
// 	UpdatedAt           time.Time          `json:"updated_at"`
// 	ApprovedBy          int64              `json:"approved_by"`
// 	InvoiceDate         time.Time          `json:"invoice_date"`
// 	Observations        string             `json:"observations"`
// 	ProratedDays        int64              `json:"prorated_days"`
// 	PaymentMethod       string             `json:"payment_method"`
// 	AlegraTransaction   interface{}        `json:"alegra_transaction"`
// 	SubscriptionOrderID int64              `json:"subscription_order_id"`
// }

// type DetailListResponse struct {
// 	TotalsListResponse TotalsListResponse `json:"totals"`
// 	// CreatedAtUser                  time.Time                      `json:"created_at_user"`
// 	// FirstSaleDate                  time.Time                      `json:"first_sale_date"`
// 	// GroupCreatedAt                 time.Time                      `json:"group_created_at"`
// 	// IsFirstPurchase                bool                           `json:"is_first_purchase"`
// 	// DetailsBeforeFirstPurchaseList DetailsBeforeFirstPurchaseList `json:"details_before_first_purchase"`
// 	// DaysElapsedPreviousPurchase    int64                          `json:"days_elapsed_previous_purchase"`
// }

// type TotalsListResponse struct {
// 	Money   float64 `json:"money"`
// 	Tickets int64   `json:"tickets"`
// }

// type DetailsBeforeFirstPurchaseListResponse struct {
// 	DaysElapsed   int64 `json:"days_elapsed"`
// 	SubjectsViews int64 `json:"subjects_views"`
// }

type InvoiceItemListResponse struct {
	// ID                  int64     `json:"id"`
	// Quantity            int64     `json:"quantity"`
	// CourseID            int64     `json:"course_id"`
	// BasePrice           int64     `json:"base_price"`
	// CreatedAt           time.Time `json:"created_at"`
	// InvoiceID           int64     `json:"invoice_id"`
	// UnitPrice           int64     `json:"unit_price"`
	// UpdatedAt           time.Time `json:"updated_at"`
	// CurrencyID          int64     `json:"currency_id"`
	// ExchangeRate        int64     `json:"exchange_rate"`
	// DiscountPrice       int64     `json:"discount_price"`
	OriginalPrice float64 `json:"original_price"`
	// IsSubscription      bool      `json:"is_subscription"`
	// IsAttendingCourse   bool      `json:"is_attending_course"`
	// ProfessorPercentage float64   `json:"professor_percentage"`
}
