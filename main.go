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
	word := "first"
	definition := "First Word"

	d := mydict.Dictionary{word: definition}
	definition, err := d.Search(word)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(definition)
	}

	word = "second"
	definition = "Second Word"

	err = d.Add(word, definition)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(d)
	}

	err = d.Update(word, "First Word")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(d)
	}

	d.Delete(word)
	fmt.Println(d)
}
