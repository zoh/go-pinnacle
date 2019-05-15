package endpoints

import (
	"go-pinnacle/models"
	"go-pinnacle/request"
)

type Account struct {
	request.Request
}

func NewAccount(r request.Request) *Account {
	return &Account{Request: r}
}

/*
   Get Account info.
   :return: client balance
*/
func (a *Account) GetClientBalance() (*models.ClientBalanceResponse, error) {
	var res = new(models.ClientBalanceResponse)
	err := a.Send(res, "GET", "v1/client/balance", nil, nil)
	return res, err
}
