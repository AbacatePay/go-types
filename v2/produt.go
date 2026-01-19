package v2

import "time"

// https://docs.abacatepay.com/pages/products/reference#estrutura
type APIProduct struct {
	// The ID of your product.
	ID string `json:"id"`

	// Unique product identifier in your system.
	ExternalID string `json:"externalId"`

	// Product name.
	Name string `json:"name"`

	// Product price in cents.
	Price uint64 `json:"price"`

	// Product currency.
	Currency string `json:"currency"`

	// Product status.
	Status ProductStatus `json:"status"`

	// Indicates whether the product was created in a testing environment.
	DevMode bool `json:"devMode"`

	// Product creation date.
	CreatedAt time.Time `json:"createdAt"`

	// Product update date.
	UpdatedAt time.Time `json:"updatedAt"`

	// Product description.
	Description *string `json:"description,omitempty"`
}

type ProductStatus string

const (
	ProductStatusActive   ProductStatus = "ACTIVE"
	ProductStatusInactive ProductStatus = "INACTIVE"
)
