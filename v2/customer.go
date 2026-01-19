package v2

// https://docs.abacatepay.com/pages/client/reference#estrutura
type APICustomer struct {
	// Unique customer identifier.
	ID string `json:"id"`

	// Indicates whether the client was created in a testing environment.
	DevMode bool `json:"devMode"`

	// Customer country.
	Country string `json:"country"`

	// Customer's full name.
	Name string `json:"name"`

	// Customer's email.
	Email string `json:"email"`

	// Customer's CPF or CNPJ.
	TaxID string `json:"taxId"`

	// Customer's cell phone.
	Cellphone string `json:"cellphone"`

	// Customer zip code.
	ZipCode string `json:"zipCode"`

	// Additional customer metadata.
	Metadata map[string]any `json:"metadata,omitempty"`
}

type APIBaseCustomer struct {
	// Customer's full name.
	Name string `json:"name"`

	// Customer's email.
	Email string `json:"email"`

	// Customer's CPF or CNPJ.
	TaxID string `json:"taxId"`

	// Customer's cell phone.
	Cellphone string `json:"cellphone"`
}
