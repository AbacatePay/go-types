package v1

// https://docs.abacatepay.com/pages/trustMRR/reference#estrutura
type APIMerchantInfo struct {
	Name      string  `json:"name"`
	Website   *string `json:"website"`
	CreatedAt string  `json:"createdAt"`
}

// https://docs.abacatepay.com/pages/trustMRR/reference#mrr-monthly-recurring-revenue
type APIMRRInfo struct {
	MRR                      uint64 `json:"mrr"`
	TotalActiveSubscriptions uint64 `json:"totalActiveSubscriptions"`
}
