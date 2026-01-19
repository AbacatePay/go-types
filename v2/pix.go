package v2

import "time"

// https://docs.abacatepay.com/pages/transparents/reference
type APIQRCodePIX struct {
	// Unique QRCode PIX identifier.
	ID string

	// Charge amount in cents (e.g. 4000 = R$40.00).
	Amount uint64

	// PIX status. Can be `PENDING`, `EXPIRED`, `CANCELLED`, `PAID`, `REFUNDED`.
	Status PaymentStatus

	// Indicates whether the charge is in a testing (true) or production (false) environment.
	DevMode bool

	// PIX code (copy-and-paste) for payment.
	BrCode string

	// PIX code in Base64 format (Useful for displaying in images).
	BrCodeBase64 string

	// Platform fee in cents. Example: 80 means R$0.80.
	PlatformFee uint64

	// QRCode PIX creation date and time.
	CreatedAt time.Time

	// QRCode PIX last updated date and time.
	UpdatedAt time.Time

	// QRCode expiration date and time.
	ExpiresAt time.Time
}
