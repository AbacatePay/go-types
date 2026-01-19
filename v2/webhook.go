package v2

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
)

const AbacatePayPublicKey = "t9dXRhHHo3yDEj5pVDYz0frf7q6bMKyMRmxxCPIPp3RCplBfXRxqlC6ZpiWmOqj4L63qEaeUOtrCI8P0VMUgo6iIga2ri9ogaHFs0WIIywSMg0q7RmBfybe1E5XJcfC4IW3alNqym0tXoAKkzvfEjZxV6bE0oG2zJrNNYmUCKZyV0KZ3JS8Votf9EAWWYdiDkMkpbMdPggfh1EqHlVkMiTady6jOR3hyzGEHrIz2Ret0xHKMbiqkr9HS1JhNHDX9"

// VerifyWebhookSignature verifies whether the webhook signature is valid.
func VerifyWebhookSignature(rawBody string, signature string) bool {
	mac := hmac.New(sha256.New, []byte(AbacatePayPublicKey))
	mac.Write([]byte(rawBody))

	expectedSig := mac.Sum(nil)

	receivedSig, err := base64.StdEncoding.DecodeString(signature)

	if err != nil {
		return false
	}

	return subtle.ConstantTimeCompare(expectedSig, receivedSig) == 1
}

// https://docs.abacatepay.com/pages/webhooks
type WebhookEventType string

// https://docs.abacatepay.com/pages/webhooks
const (
	WebhookEventPayoutFailed WebhookEventType = "payout.failed"
	WebhookEventPayoutDone   WebhookEventType = "payout.done"
	WebhookEventBillingPaid  WebhookEventType = "billing.paid"
)

type BaseWebhookEvent[Paylod any] struct {
	// Unique identifier for the webhook.
	ID string `json:"id"`

	// Event type.
	Event WebhookEventType `json:"event"`

	// Event data.
	Data Paylod `json:"data"`

	// Indicates whether the event occurred in dev mode.
	DevMode bool `json:"devMode"`
}

// https://docs.abacatepay.com/pages/webhooks#payout-failed
type WebhookPayoutFailedEvent BaseWebhookEvent[WebhookPayoutFailedEventData]

// https://docs.abacatepay.com/pages/webhooks#payout-failed
type WebhookPayoutFailedEventData struct {
	// Status is always CANCELLED.
	Transaction APIPayout `json:"transaction"`
}

// https://docs.abacatepay.com/pages/webhooks#payout-done
type WebhookPayoutDoneEventData struct {
	Transaction APIPayout `json:"transaction"`
}

// https://docs.abacatepay.com/pages/webhooks#payout-done
type WebhookPayoutDoneEvent BaseWebhookEvent[WebhookPayoutDoneEventData]

type WebhookBillingPaidPayment struct {
	Fee    uint64        `json:"fee"`
	Method PaymentMethod `json:"method"`
	Amount uint64        `json:"amount"`
}

type WebhookBillingPaidPix struct {
	// Charge amount in cents (e.g. 4000 = R$40.00).
	Amount uint64 `json:"amount"`

	// Unique billing identifier.
	ID string `json:"id"`

	// Always PIX
	Kind PaymentMethod `json:"kind"`

	// Billing status, can only be `PAID` here.
	Status PaymentStatus `json:"status"`
}

type WebhookBillingPaidBilling struct {
	// Charge amount in cents (e.g. 4000 = R$40.00).
	Amount int `json:"amount"`

	// Unique billing identifier.
	ID string `json:"id"`

	// Bill ID in your system.
	ExternalID string `json:"externalId"`

	// Status of the payment. Always `PaymentStatusPaid`.
	Status PaymentStatus `json:"status"`

	// URL where the user can complete the payment.
	URL string `json:"url"`
}

// https://docs.abacatepay.com/pages/webhooks#billing-paid
type WebhookBillingPaidEventData struct {
	Payment WebhookBillingPaidPayment `json:"payment"`

	// Exactly one of these will be non-nil

	PixQrCode *WebhookBillingPaidPix     `json:"pixQrCode,omitempty"`
	Billing   *WebhookBillingPaidBilling `json:"billing,omitempty"`
}

// // https://docs.abacatepay.com/pages/webhooks#billing-paid
type WebhookBillingPaidEvent BaseWebhookEvent[WebhookBillingPaidEventData]

type WebhookEnvelope struct {
	Event WebhookEventType `json:"event"`
}

func (event WebhookEnvelope) IsPayoutFailed() bool {
	return event.Event == WebhookEventPayoutFailed
}

func (event WebhookEnvelope) IsPayoutDone() bool {
	return event.Event == WebhookEventPayoutDone
}

func (event WebhookEnvelope) IsBillingPaid() bool {
	return event.Event == WebhookEventBillingPaid
}

func (event WebhookBillingPaidEvent) IsPix() bool {
	return event.Data.PixQrCode != nil
}

func (event WebhookBillingPaidEvent) IsBilling() bool {
	return event.Data.Billing != nil
}
