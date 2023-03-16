package view

import (
	"github.com/gin-gonic/gin"
	"github.com/vaberof/it-planet-task/internal/domain/account"
)

type AccountResponse struct {
	Id        int32  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName" `
	Email     string `json:"email"`
}

func RenderAccountResponse(c *gin.Context, statusCode int, account *account.Account) {
	RenderResponse(c, statusCode, buildAccount(account))
}

func RenderAccountsResponse(c *gin.Context, statusCode int, accounts []*account.Account) {
	accountsResponse := make([]*AccountResponse, len(accounts))

	for i := 0; i < len(accountsResponse); i++ {
		accountsResponse[i] = buildAccount(accounts[i])
	}

	RenderResponse(c, statusCode, accountsResponse)
}

func buildAccount(account *account.Account) *AccountResponse {
	return &AccountResponse{
		Id:        account.Id,
		FirstName: account.FirstName,
		LastName:  account.LastName,
		Email:     account.Email,
	}
}
