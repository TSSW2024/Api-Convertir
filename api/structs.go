package api

type ConversionRequest struct {
	Amount string `json:"amount"`
	From   string `json:"from"`
	To     string `json:"to"`
}

type ConversionResponse struct {
	Amount string  `json:"amount"`
	From   string  `json:"from"`
	To     string  `json:"to"`
	Result float64 `json:"result"`
}
