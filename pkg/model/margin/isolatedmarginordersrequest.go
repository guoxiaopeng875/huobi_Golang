package margin

type IsolatedMarginOrdersRequest struct {
	Symbol   string `json:"symbol"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}
