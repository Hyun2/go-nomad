package main

import (
	"fmt"

	"github.com/hyun2/nomad-learn-go/account"
)

func main() {
	account := account.MakeAccount("Stoker")
	account.Deposit(10)
	// err := account.Withdraw(20)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(account.Balance())
	fmt.Println(account)
}
