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

type Result struct {
	url    string
	status string
}

func main() {
	ch := make(chan Result)
	results := map[string]string{}

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
		go hitUrl(url, ch)
	}

	for i := 0; i < len(urls); i++ {
		result := <-ch
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitUrl(url string, ch chan<- Result) {
	// fmt.Println("Checking Url:", url)

	resp, err := http.Get(url)
	status := "OK"

	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}

	ch <- Result{url: url, status: status}
}
