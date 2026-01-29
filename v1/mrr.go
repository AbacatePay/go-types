package v1

// https://docs.abacatepay.com/pages/trustMRR/reference#estrutura
type APIMerchantInfo struct {
	Name      string  `json:"name"`
	Website   *string `json:"website"`
	CreatedAt string  `json:"createdAt"`
}
