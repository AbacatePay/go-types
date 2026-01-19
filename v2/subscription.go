package v2

import "time"

// https://docs.abacatepay.com/pages/subscriptions/reference#estrutura
type APISubscription struct {
	// The ID of the subscription.
	ID string

	// The subscription value in cents.
	Amount uint64

	// Subscription currency.
	Currency string

	// Subscription name.
	Name string

	// Subscription description.
	Description string

	// Unique identifier of the subscription on your system.
	ExternalID string

	// Indicates whether the signature was created in a testing environment.
	DevMode bool

	// Subscription creation date.
	CreatedAt time.Time

	// Subscription update date.
	UpdatedAt time.Time

	// Payment method for the subscription.
	Method PaymentMethod

	// Status of the subscription.
	Status SubscriptionStatus

	// Billing frequency configuration.
	Frequency APISubscriptionFrequency

	// Identifier of the customer who will have the signature.
	CustomerID string

	// Retry policy in case of payment failure.
	RetryPolicy APISubscriptionRetryPolicy

	// Array of events related to the subscription.
	Events []APISubscriptionEvent
}

type APISubscriptionFrequency struct {
	// Subscription billing cycle.
	Cycle APISubscriptionFrequencyCycle

	// Day of the month the charge will be processed (1-31).
	DayOfProcessing uint8
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
	MaxRetry uint64

	// Interval in days between charging attempts.
	RetryEvery uint64
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
	Event string

	// Event description.
	Description string

	// Event creation date.
	CreatedAt time.Time
}
