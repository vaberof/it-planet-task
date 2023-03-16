package handler

import "github.com/vaberof/it-planet-task/internal/domain/account"

type AccountService interface {
	CreateAccount(firstName, lastName, email, password string) (*account.Account, error)
	GetAccountById(id int32) (*account.Account, error)
	UpdateAccount(initialAccountId int32, initialEmail string, accountId int32, firstName, lastName, email, password string) (*account.Account, error)
	SearchAccounts(firstName, lastName, email string, from, size int) ([]*account.Account, error)
	DeleteAccount(initialAccountId, accountId int32) error
}
