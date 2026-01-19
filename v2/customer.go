package v2

// https://docs.abacatepay.com/pages/client/reference#estrutura
type APICustomer struct {
	// Unique customer identifier.
	ID string

	// Indicates whether the client was created in a testing environment.
	DevMode bool

	// Customer country.
	Country string

	// Customer's full name.
	Name string

	// Customer's email.
	Email string

	// Customer's CPF or CNPJ.
	TaxID string

	// Customer's cell phone.
	Cellphone string

	// Customer zip code.
	ZipCode string

	// Additional customer metadata.
	Metadata map[string]any
}
