package account

type Account struct {
	owner   string
	balance int
}

func MakeAccount(owner string) *Account {
	newAccount := Account{owner: owner, balance: 0}

	return &newAccount
}

func (account *Account) Deposit(amount int) {
	account.balance += amount
}

func (account *Account) Balance() int {
	return account.balance
}
