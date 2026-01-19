package v2

import "time"

// https://docs.abacatepay.com/pages/checkouts/reference#estrutura
type APICheckout struct {
	// Unique billing identifier.
	ID string `json:"id"`

	// Total amount to be paid in cents.
	Amount uint64 `json:"amount"`

	// Amount already paid in cents.
	PaidAmount *uint64 `json:"paidAmount,omitempty"`

	// Bill ID in your system.
	ExternalID *string `json:"externalId,omitempty"`

	// URL where the user can complete the payment.
	URL string `json:"url"`

	// List of items in billing.
	Items []APIItem `json:"items"`

	// Billing status. Can be `PENDING`, `EXPIRED`, `CANCELLED`, `PAID`, `REFUNDED`.
	Status PaymentStatus `json:"status"`

	// Indicates whether the charge was created in a development (true) or production (false) environment.
	DevMode bool `json:"devMode"`

	// Additial metadata for the charge.
	Metadata map[string]any `json:"metadata,omitempty"`

	// URL that the customer will be redirected to when clicking the “back” button.
	ReturnURL string `json:"returnUrl"`

	// URL that the customer will be redirected to when making payment.
	CompletionURL string `json:"completionUrl"`

	// Payment receipt URL.
	ReceiptURL *string `json:"receiptUrl,omitempty"`

	// Coupons allowed in billing.
	Coupons []string `json:"coupons"`

	// Customer ID associated with the charge.
	CustomerID *string `json:"customerId,omitempty"`

	// Charge creation date and time.
	CreatedAt time.Time `json:"createdAt"`

	// Charge last updated date and time.
	UpdatedAt time.Time `json:"updatedAt"`
}

type APIItem struct {
	// Product ID.
	ID string `json:"id"`

	// Item quantity.
	Quantity uint64 `json:"quantity"`
}

// https://docs.abacatepay.com/pages/checkouts/reference#atributos
type PaymentStatus string

// https://docs.abacatepay.com/pages/checkouts/reference#atributos
const (
	PaymentStatusPending   PaymentStatus = "PENDING"
	PaymentStatusExpired   PaymentStatus = "EXPIRED"
	PaymentStatusCancelled PaymentStatus = "CANCELLED"
	PaymentStatusPaid      PaymentStatus = "PAID"
	PaymentStatusRefunded  PaymentStatus = "REFUNDED"
)

// https://docs.abacatepay.com/pages/checkouts/create#body-methods
type PaymentMethod string

// https://docs.abacatepay.com/pages/checkouts/create#body-methods
const (
	PaymentMethodPix  PaymentMethod = "PIX"
	PaymentMethodCard PaymentMethod = "CARD"
)
