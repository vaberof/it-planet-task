package auth

import (
	"errors"
	"github.com/vaberof/it-planet-task/internal/domain/account"
	"github.com/vaberof/it-planet-task/internal/pkg/xerror"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type AuthService struct {
	accountStorage AccountStorage
}

func NewAuthService(accountStorage AccountStorage) *AuthService {
	return &AuthService{
		accountStorage: accountStorage,
	}
}

func (s *AuthService) AuthenticateAccount(r *http.Request) (*account.Account, error) {
	return s.authenticateAccountImpl(r)
}

func (s *AuthService) authenticateAccountImpl(r *http.Request) (*account.Account, error) {
	email, password, hasAuth := r.BasicAuth()
	if !hasAuth {
		return nil, xerror.NewErrorWrapper(
			401,
			"request from unauthorized account",
			errors.New("unauthorized"))
	}

	domainAccount, err := s.accountStorage.GetAccountByEmail(email)
	if err != nil {
		return nil, xerror.NewErrorWrapper(
			401,
			"incorrect authorization data",
			errors.New("incorrect authorization data"))
	}

	err = s.validatePassword(domainAccount.Password, password)
	if err != nil {
		return nil, xerror.NewErrorWrapper(
			401,
			"incorrect authorization data",
			errors.New("incorrect authorization data"))
	}

	return domainAccount, nil
}

func (s *AuthService) validatePassword(hashedPassword string, inputPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword)); err != nil {
		return err
	}
	return nil
}
