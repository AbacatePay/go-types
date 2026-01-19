package v2

import "time"

// https://docs.abacatepay.com/pages/coupon/reference#estrutura
type APICoupon struct {
	// Unique coupon code that your customers will use to apply the discount.
	ID string `json:"id"`

	// Type of discount applied by the coupon.
	DiscountKind CouponDiscountKind `json:"discountKind"`

	// Discount amount. For `PERCENTAGE` use numbers from 1-100 (ex: 10 = 10%). For `FIXED` use the value in cents (ex: 500 = R$5.00).
	Discount uint64 `json:"discount"`

	// Coupon status.
	Status CouponStatus `json:"status"`

	// Limit on the number of times the coupon can be used. Use `-1` for unlimited coupons or a specific number to limit usage.
	MaxRedeems uint64 `json:"maxRedeems"`

	// Counter of how many times the coupon has been used by customers.
	RedeemsCount uint64 `json:"redeemsCount"`

	// Indicates whether the coupon was created in a development (true) or production (false) environment.
	DevMode bool `json:"devMode"`

	// Internal description of the coupon for your organization and control.
	Notes *string `json:"notes,omitempty"`

	// Coupon creation date and time.
	CreatedAt time.Time `json:"createdAt"`

	// Coupon last updated date and time.
	UpdatedAt time.Time `json:"updatedAt"`

	// Object to store additional information about the coupon, such as campaign, category, or other data relevant to your organization.
	Metadata map[string]any `json:"metadata,omitempty"`
}

// https://docs.abacatepay.com/pages/coupon/reference#atributos
type CouponDiscountKind string

// https://docs.abacatepay.com/pages/coupon/reference#atributos
const (
	CouponDiscountKindPercentage CouponDiscountKind = "PERCENTAGE"
	CouponDiscountKindFixed      CouponDiscountKind = "FIXED"
)

// https://docs.abacatepay.com/pages/coupon/reference#atributos
type CouponStatus string

// https://docs.abacatepay.com/pages/coupon/reference#atributos
const (
	CouponStatusActive   CouponStatus = "ACTIVE"
	CouponStatusInactive CouponStatus = "INACTIVE"
	CouponStatusExpired  CouponStatus = "EXPIRED"
)
