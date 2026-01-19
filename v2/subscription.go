package v2

import "time"

// https://docs.abacatepay.com/pages/subscriptions/reference#estrutura
type APISubscription struct {
	// The ID of the subscription.
	ID string `json:"id"`

	// The subscription value in cents.
	Amount uint64 `json:"amount"`

	// Subscription currency.
	Currency string `json:"currency"`

	// Subscription name.
	Name string `json:"name"`

	// Subscription description.
	Description string `json:"description"`

	// Unique identifier of the subscription on your system.
	ExternalID string `json:"externalId"`

	// Indicates whether the signature was created in a testing environment.
	DevMode bool `json:"devMode"`

	// Subscription creation date.
	CreatedAt time.Time `json:"createdAt"`

	// Subscription update date.
	UpdatedAt time.Time `json:"updatedAt"`

	// Payment method for the subscription.
	Method PaymentMethod `json:"method"`

	// Status of the subscription.
	Status SubscriptionStatus `json:"status"`

	// Billing frequency configuration.
	Frequency APISubscriptionFrequency `json:"frequency"`

	// Identifier of the customer who will have the signature.
	CustomerID string `json:"customerId"`

	// Retry policy in case of payment failure.
	RetryPolicy APISubscriptionRetryPolicy `json:"retryPolicy"`

	// Array of events related to the subscription.
	Events []APISubscriptionEvent `json:"events"`
}

type APISubscriptionFrequency struct {
	// Subscription billing cycle.
	Cycle APISubscriptionFrequencyCycle `json:"cycle"`

	// Day of the month the charge will be processed (1-31).
	DayOfProcessing uint8 `json:"dayOfProcessing"`
}

type APISubscriptionFrequencyCycle string

const (
	APISubscriptionFrequencyCycleMonthly APISubscriptionFrequencyCycle = "MONTHLY"
	APISubscriptionFrequencyCycleYearly  APISubscriptionFrequencyCycle = "YEARLY"
	APISubscriptionFrequencyCycleWeekly  APISubscriptionFrequencyCycle = "WEEKLY"
	APISubscriptionFrequencyCycleDaily   APISubscriptionFrequencyCycle = "DAILY"
)

type APISubscriptionRetryPolicy struct {
	// Maximum number of billing attempts.
	MaxRetry uint64 `json:"maxRetry"`

	// Interval in days between charging attempts.
	RetryEvery uint64 `json:"retryEvery"`
}

type SubscriptionStatus string

const (
	SubscriptionStatusPending   SubscriptionStatus = "PENDING"
	SubscriptionStatusActive    SubscriptionStatus = "ACTIVE"
	SubscriptionStatusCancelled SubscriptionStatus = "CANCELLED"
	SubscriptionStatusExpired   SubscriptionStatus = "EXPIRED"
	SubscriptionStatusFailed    SubscriptionStatus = "FAILED"
)

// https://docs.abacatepay.com/pages/subscriptions/reference#estrutura
type APISubscriptionEvent struct {
	// Event type.
	Event string `json:"event"`

	// Event description.
	Description string `json:"description"`

	// Event creation date.
	CreatedAt time.Time `json:"createdAt"`
}
