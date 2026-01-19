package v2

// https://docs.abacatepay.com/pages/store/reference#estrutura
type APIStore struct {
	// Unique identifier for your store on AbacatePay.
	ID string `json:"id"`

	// Name of your store/company.
	Name string `json:"name"`

	// Object containing information about your account balances.
	// All balance values ​​are returned in cents. To convert to Reais, divide by 100. For example: 15000 cents = R$150.00.
	Balance APIStoreBalance `json:"balance"`
}

// ps://docs.abacatepay.com/pages/store/reference#estrutura
type APIStoreBalance struct {
	// Balance available for withdrawal in cents.
	Available uint16 `json:"available"`

	// Balance pending confirmation in cents.
	Pending uint16 `json:"pending"`

	// Balance blocked in disputes in cents.
	// The blocked balance represents amounts that are in dispute or under review. These amounts are not available for withdrawal until the situation is resolved.
	Blocked uint16 `json:"blocked"`
}
