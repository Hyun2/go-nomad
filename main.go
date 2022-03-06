package main

import (
	"errors"
	"fmt"
	"net/http"
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
// func main() {
// 	word := "first"
// 	definition := "First Word"

// 	d := mydict.Dictionary{word: definition}
// 	definition, err := d.Search(word)
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		fmt.Println(definition)
// 	}

// 	word = "second"
// 	definition = "Second Word"

// 	err = d.Add(word, definition)
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		fmt.Println(d)
// 	}

// 	err = d.Update(word, "First Word")
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		fmt.Println(d)
// 	}

// 	d.Delete(word)
// 	fmt.Println(d)
// }

// hitUrl

var errRequestFailed = errors.New("request Failed")

func main() {
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		hitUrl(url)
	}
}

func hitUrl(url string) error {
	fmt.Println("Checking Url:", url)

	resp, err := http.Get(url)

	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
