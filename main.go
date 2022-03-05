package main

import (
	"fmt"

	"github.com/hyun2/nomad-learn-go/account"
)

func main() {
	account := account.MakeAccount("Stoker")
	account.Deposit(10)
	fmt.Println(account.Balance())
}
