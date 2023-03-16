package account

type AccountStorage interface {
	CreateAccount(firstName, lastName, email, password string) (*Account, error)
	GetAccountByEmail(email string) (*Account, error)
	GetAccountById(id int32) (*Account, error)
	UpdateAccount(id int32, firstName, lastName, email, password string) (*Account, error)
	FindAccountByEmail(email string) error
	GetFilteredAccounts(firstName, lastName, email string, from, size int) ([]*Account, error)
	DeleteAccount(id int32) error
	IsAccountAssociatedWithAnimal(accountId int32) error
}
