package account

import (
	"errors"
	"fmt"
)

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

func (account Account) Balance() int {
	return account.balance
}

func (account *Account) Withdraw(amount int) error {
	if account.balance < amount {
		return errors.New("can't withdraw, you're poor")
	}
	account.balance -= amount
	return nil
}

func (account *Account) ChangeOwner(newOwner string) {
	account.owner = newOwner
}

func (account Account) Owner() string {
	return account.owner
}

func (account Account) String() string {
	return fmt.Sprintf("Owner: %s\nAmount: %d", account.owner, account.balance)
}
