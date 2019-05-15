package go_pinnacle

import (
	"go-pinnacle/endpoints"
	"go-pinnacle/request"
)

type ApiClient struct {
	request request.Request
	Account *endpoints.Account
	Lines   *endpoints.Lines
	Bets    *endpoints.Bets
}

func NewApiClient(username, password string) *ApiClient {
	r := request.NewRequest(username, password, "https://api.pinnacle.com/", nil)
	api := &ApiClient{request: r}

	api.Account = endpoints.NewAccount(r)
	api.Lines = endpoints.NewLines(r)
	api.Bets = endpoints.NewBets(r)

	return api
}
