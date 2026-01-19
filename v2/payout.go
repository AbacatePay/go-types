package v2

import "time"

// https://docs.abacatepay.com/pages/payouts/reference
type APIPayout struct {
	// Unique transaction identifier.
	ID string `json:"id"`

	// Current transaction status.
	Status PayoutStatus `json:"status"`

	// Indicates whether the transaction was created in a testing environment.
	DevMode bool `json:"devMode"`

	// Transaction proof URL.
	ReceiptURL *string `json:"receiptUrl,omitempty"`

	// Payout value in cents.
	Amount uint64 `json:"amount"`

	// Platform fee in cents.
	PlatformFee uint64 `json:"platformFee"`

	// External transaction identifier.
	ExternalID string `json:"externalId"`

	// Transaction creation date.
	CreatedAt time.Time `json:"createdAt"`

	// Transaction update date.
	UpdatedAt time.Time `json:"updatedAt"`
}

// https://docs.abacatepay.com/pages/payouts/reference#atributos
type PayoutStatus string

// https://docs.abacatepay.com/pages/payouts/reference#atributos
const (
	PayoutStatusPending   PayoutStatus = "PENDING"
	PayoutStatusExpired   PayoutStatus = "EXPIRED"
	PayoutStatusCancelled PayoutStatus = "CANCELLED"
	PayoutStatusComplete  PayoutStatus = "COMPLETE"
	PayoutStatusRefunded  PayoutStatus = "REFUNDED"
)
