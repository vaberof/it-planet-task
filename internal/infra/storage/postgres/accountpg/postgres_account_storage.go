package accountpg

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/vaberof/it-planet-task/internal/domain/account"
	"github.com/vaberof/it-planet-task/internal/infra/storage/postgres/animalpg"
	"github.com/vaberof/it-planet-task/internal/pkg/xerror"
	"gorm.io/gorm"
	"strings"
)

type PostgresAccountStorage struct {
	db *gorm.DB
}

func NewPostgresAccountStorage(db *gorm.DB) *PostgresAccountStorage {
	return &PostgresAccountStorage{
		db: db,
	}
}

func (s *PostgresAccountStorage) CreateAccount(firstName, lastName, email, password string) (*account.Account, error) {
	var postgresAccount Account

	postgresAccount.FirstName = firstName
	postgresAccount.LastName = lastName
	postgresAccount.Email = email
	postgresAccount.Password = password

	err := s.db.Table("accounts").Create(&postgresAccount).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "accountpg",
			"func":    "CreateAccount",
		}).Error(err)

		return nil, xerror.NewErrorWrapper(500, "cannot create account", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "accountpg",
		"func":    "CreateAccount",
	}).Info("created account")

	return BuildDomainAccount(&postgresAccount), nil
}

func (s *PostgresAccountStorage) GetAccountById(id int32) (*account.Account, error) {
	postgresAccount, err := s.getAccountById(id)
	if err != nil {
		return nil, err
	}

	return BuildDomainAccount(postgresAccount), nil
}

func (s *PostgresAccountStorage) GetAccountByEmail(email string) (*account.Account, error) {
	postgresAccount, err := s.getAccountByEmail(email)
	if err != nil {
		return nil, err
	}

	return BuildDomainAccount(postgresAccount), nil
}

func (s *PostgresAccountStorage) UpdateAccount(id int32, firstName, lastName, email, password string) (*account.Account, error) {
	postgresAccount, err := s.getAccountById(id)
	if err != nil {
		return nil, xerror.NewErrorWrapper(403, "cannot find account with given id", err)
	}

	postgresAccount.FirstName = firstName
	postgresAccount.LastName = lastName
	postgresAccount.Email = email
	postgresAccount.Password = password

	err = s.db.Table("accounts").Save(&postgresAccount).Error
	if err != nil {
		return nil, xerror.NewErrorWrapper(500, "cannot update account", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "accountpg",
		"func":    "UpdateAccount",
	}).Info("updated account")

	return BuildDomainAccount(postgresAccount), nil
}

func (s *PostgresAccountStorage) GetFilteredAccounts(firstName, lastName, email string, from, size int) ([]*account.Account, error) {
	var postgresAccounts []*Account
	var query strings.Builder

	query.WriteString("1=1")

	if firstName != "" {
		query.WriteString(fmt.Sprint(" AND LOWER(first_name) LIKE ", "LOWER('%"+firstName+"%')"))
	}

	if lastName != "" {
		query.WriteString(fmt.Sprint(" AND LOWER(last_name) LIKE ", "LOWER('%"+lastName+"%')"))
	}

	if email != "" {
		query.WriteString(fmt.Sprint(" AND LOWER(email) LIKE ", "LOWER('%"+email+"%')"))
	}

	err := s.db.
		Table("accounts").
		Where(query.String()).
		Order("id ASC").
		Offset(from).
		Limit(size).
		Find(&postgresAccounts).Error

	if err != nil {
		return nil, xerror.NewErrorWrapper(500, "cannot find filtered accounts", err)
	}

	return BuildDomainAccounts(postgresAccounts), nil
}

func (s *PostgresAccountStorage) DeleteAccount(id int32) error {
	postgresAccount, err := s.getAccountById(id)
	if err != nil {
		return xerror.NewErrorWrapper(403, "cannot find account with given id", err)
	}

	err = s.db.Table("accounts").Where("id = ?", id).Delete(&postgresAccount).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "accountpg",
			"func":    "DeleteAccount",
		}).Error(err)

		return xerror.NewErrorWrapper(500, "cannot delete account", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "accountpg",
		"func":    "DeleteAccount",
	}).Info("deleted account")

	return nil
}

func (s *PostgresAccountStorage) FindAccountById(id int32) error {
	var postgresAccount Account

	err := s.db.Table("accounts").Where("id = ?", id).First(&postgresAccount).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "accountpg",
			"func":    "FindAccountById",
		}).Error(err)

		return xerror.NewErrorWrapper(404, "cannot find account with given id", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "accountpg",
		"func":    "FindAccountById",
	}).Info("found account by id")

	return nil
}

func (s *PostgresAccountStorage) FindAccountByEmail(email string) error {
	var postgresAccount Account

	err := s.db.Table("accounts").Where("email = ?", email).First(&postgresAccount).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "accountpg",
			"func":    "FindAccountByEmail",
		}).Error(err)

		return xerror.NewErrorWrapper(404, "cannot find account with given email", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "accountpg",
		"func":    "FindAccountByEmail",
	}).Info("found account by email")

	return nil
}

func (s *PostgresAccountStorage) IsAccountAssociatedWithAnimal(accountId int32) error {
	var postgresAnimal animalpg.Animal

	err := s.db.Table("animals").Where("chipper_id = ?", accountId).First(&postgresAnimal).Error
	if err == nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "accountpg",
			"func":    "IsAccountAssociatedWithAnimal",
		}).Info("account associated with animal")

		return xerror.NewErrorWrapper(
			400,
			fmt.Sprintf("account associated with animal with %d id", postgresAnimal.Id),
			errors.New("cannot delete associated account"))
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "accountpg",
		"func":    "IsAccountAssociatedWithAnimal",
	}).Info("account is not associated with any of animals")

	return nil
}

func (s *PostgresAccountStorage) getAccountById(id int32) (*Account, error) {
	var postgresAccount Account

	err := s.db.Table("accounts").Where("id = ?", id).First(&postgresAccount).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "accountpg",
			"func":    "getAccountById",
		}).Error(err)

		return nil, xerror.NewErrorWrapper(404, "cannot find account with given id", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "accountpg",
		"func":    "getAccountById",
	}).Info("received account by id")

	return &postgresAccount, nil
}

func (s *PostgresAccountStorage) getAccountByEmail(email string) (*Account, error) {
	var postgresAccount Account

	err := s.db.Table("accounts").Where("email = ?", email).First(&postgresAccount).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "accountpg",
			"func":    "getAccountByEmail",
		}).Error(err)

		return nil, xerror.NewErrorWrapper(404, "cannot find account with given email", err)
	}

	logrus.WithFields(logrus.Fields{
		"layer":   "infra",
		"package": "accountpg",
		"func":    "getAccountByEmail",
	}).Info("received account by email")

	return &postgresAccount, nil
}
