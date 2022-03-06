package main

import (
	"fmt"
	"log"

	"github.com/hyun2/nomad-learn-go/mydict"
)

// account example
// func main() {
// 	account := account.MakeAccount("Stoker")
// 	account.Deposit(10)
// 	// err := account.Withdraw(20)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// fmt.Println(account.Balance())
// 	fmt.Println(account)
// }

// mydict example
func main() {
	d := mydict.Dictionary{"first": "First Word"}
	definition, err := d.Search("second")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(definition)
	}
}
