package auth

import "github.com/vaberof/it-planet-task/internal/domain/account"

type AccountStorage interface {
	GetAccountByEmail(email string) (*account.Account, error)
}
