package v2

import "time"

// https://docs.abacatepay.com/pages/checkouts/reference#estrutura
type APICheckout struct {
	// Unique billing identifier.
	ID string

	// Total amount to be paid in cents.
	Amount	uint64

	// Amount already paid in cents.
	PaidAmount *uint64

	// Bill ID in your system.
	ExternalID *string

	// URL where the user can complete the payment.
	URL	string

	// List of items in billing.
	Items []APIItem

	// Billing status. Can be `PENDING`, `EXPIRED`, `CANCELLED`, `PAID`, `REFUNDED`.
	Status PaymentStatus

	// Indicates whether the charge was created in a development (true) or production (false) environment.
	DevMode bool

	// Additial metadata for the charge.
	Metadata map[string]any

	// URL that the customer will be redirected to when clicking the “back” button.
	ReturnURL string

	// URL that the customer will be redirected to when making payment.
	CompletionURL string

	// Payment receipt URL.
	ReceiptURL *string

	// Coupons allowed in billing.
	Coupons []string

	// Customer ID associated with the charge.
	CustomerID *string

	// Charge creation date and time.
	CreatedAt	time.Time

	// Charge last updated date and time.
	UpdatedAt	time.Time
}

type APIItem struct {
	// Product ID.
	ID	string

	// Item quantity.
	Quantity	uint64
}

// https://docs.abacatepay.com/pages/checkouts/reference#atributos
type PaymentStatus string

// https://docs.abacatepay.com/pages/checkouts/reference#atributos
const (
	PaymentStatusPending PaymentStatus = "PENDING"
	PaymentStatusExpired PaymentStatus = "EXPIRED"
	PaymentStatusCancelled PaymentStatus = "CANCELLED"
	PaymentStatusPaid PaymentStatus = "PAID"
	PaymentStatusRefunded PaymentStatus = "REFUNDED"
)

// https://docs.abacatepay.com/pages/checkouts/create#body-methods
type PaymentMethod string

// https://docs.abacatepay.com/pages/checkouts/create#body-methods
const (
	PaymentMethodPix PaymentMethod = "PIX"
	PaymentMethodCard PaymentMethod = "CARD"
)
