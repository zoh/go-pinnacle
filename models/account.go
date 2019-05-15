package models

type ClientBalanceResponse struct {
	AvailableBalance        float32 `json:"availableBalance"`
	OutstandingTransactions float32 `json:"outstandingTransactions"`
	GivenCredit             float32 `json:"givenCredit"`
	Currency                string  `json:"currency"`
}
