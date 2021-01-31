package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"https://facebook.com",
		"https://amazon.com",
		"https://golang.org",
		"https://google.com",
		"https://stackoverflow.com",
	}

	for _, link := range links {
		checkLink(link)

	}

}

func checkLink(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Printf("%s is down\n", link)

	}
	fmt.Printf("%s is up\n", link)
}
