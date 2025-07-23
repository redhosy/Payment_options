package model

// PaymentOption stores payment data including account information, status, balance, and icon
type PaymentOption struct {
	Account string `json:"account"`
	Status  string `json:"status"`
	Balance string `json:"balance"`
	Icon    string `json:"icon"`
}

// Response represents the API response format
type Response struct {
	ReturnCode string                    `json:"returnCode"`
	ReturnDesc string                    `json:"returnDesc"`
	Data       map[string]*PaymentOption `json:"data"`
}
