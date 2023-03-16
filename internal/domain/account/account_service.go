package account

import (
	"errors"
	"github.com/vaberof/it-planet-task/internal/pkg/xerror"
	"golang.org/x/crypto/bcrypt"
)

type AccountService struct {
	accountStorage AccountStorage
}

func NewAccountService(accountStorage AccountStorage) *AccountService {
	return &AccountService{
		accountStorage: accountStorage,
	}
}

func (s *AccountService) CreateAccount(firstName, lastName, email, password string) (*Account, error) {
	return s.createAccountImpl(firstName, lastName, email, password)
}

func (s *AccountService) GetAccountById(id int32) (*Account, error) {
	return s.getAccountByIdImpl(id)
}

func (s *AccountService) UpdateAccount(
	initialAccountId int32,
	initialEmail string,
	accountId int32,
	firstName,
	lastName,
	email,
	password string) (*Account, error) {

	return s.updateAccountImpl(
		initialAccountId,
		initialEmail,
		accountId,
		firstName,
		lastName,
		email,
		password)
}

func (s *AccountService) SearchAccounts(firstName, lastName, email string, from, size int) ([]*Account, error) {
	return s.searchAccountsImpl(firstName, lastName, email, from, size)
}

func (s *AccountService) DeleteAccount(initialAccount, accountId int32) error {
	return s.deleteAccountImpl(initialAccount, accountId)
}

func (s *AccountService) createAccountImpl(firstName, lastName, email, password string) (*Account, error) {
	err := s.accountStorage.FindAccountByEmail(email)
	if err == nil {
		return nil, xerror.NewErrorWrapper(409,
			"account with this email already exists",
			errors.New("account with this email already exists"))
	}

	hashedPassword, err := s.hashPassword(password)
	if err != nil {
		return nil, xerror.NewErrorWrapper(500,
			"cannot create account",
			errors.New("error occurred while hashing password"))
	}

	return s.accountStorage.CreateAccount(firstName, lastName, email, hashedPassword)
}

func (s *AccountService) getAccountByIdImpl(id int32) (*Account, error) {
	return s.accountStorage.GetAccountById(id)
}

func (s *AccountService) updateAccountImpl(
	initialAccountId int32,
	initialEmail string,
	accountId int32,
	firstName,
	lastName,
	email,
	password string) (*Account, error) {

	if initialAccountId != accountId {
		return nil, xerror.NewErrorWrapper(
			403,
			"trying to update account that not belong to you",
			errors.New("invalid account id"))
	}

	if initialEmail != email {
		err := s.accountStorage.FindAccountByEmail(email)
		if err == nil {
			return nil, xerror.NewErrorWrapper(
				409,
				"account with this email already exists",
				errors.New("account with this email already exists"))
		}
	}

	hashedPassword, err := s.hashPassword(password)
	if err != nil {
		return nil, xerror.NewErrorWrapper(500, "cannot update account", err)
	}

	return s.accountStorage.UpdateAccount(accountId, firstName, lastName, email, hashedPassword)
}

func (s *AccountService) searchAccountsImpl(firstName, lastName, email string, from, size int) ([]*Account, error) {
	return s.accountStorage.GetFilteredAccounts(firstName, lastName, email, from, size)
}

func (s *AccountService) deleteAccountImpl(initialAccountId, accountId int32) error {
	if initialAccountId != accountId {
		return xerror.NewErrorWrapper(
			403,
			"trying to delete account that not belong to you",
			errors.New("invalid account id"))
	}

	err := s.accountStorage.IsAccountAssociatedWithAnimal(accountId)
	if err != nil {
		return err
	}

	return s.accountStorage.DeleteAccount(accountId)
}

func (s *AccountService) hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
