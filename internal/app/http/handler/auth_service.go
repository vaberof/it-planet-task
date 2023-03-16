package handler

import (
	"github.com/vaberof/it-planet-task/internal/domain/account"
	"net/http"
)

type AuthService interface {
	AuthenticateAccount(r *http.Request) (*account.Account, error)
}
