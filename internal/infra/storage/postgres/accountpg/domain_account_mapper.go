package accountpg

import (
	"github.com/vaberof/it-planet-task/internal/domain/account"
)

func BuildDomainAccount(postgresAccount *Account) *account.Account {
	return &account.Account{
		Id:        postgresAccount.Id,
		FirstName: postgresAccount.FirstName,
		LastName:  postgresAccount.LastName,
		Email:     postgresAccount.Email,
		Password:  postgresAccount.Password,
	}
}

func BuildDomainAccounts(postgresAccounts []*Account) []*account.Account {
	domainAccounts := make([]*account.Account, len(postgresAccounts))

	for i := 0; i < len(domainAccounts); i++ {
		domainAccounts[i] = BuildDomainAccount(postgresAccounts[i])
	}

	return domainAccounts
}
