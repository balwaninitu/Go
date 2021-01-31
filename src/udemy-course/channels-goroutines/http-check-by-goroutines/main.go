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
		go checkLink(link)

	}

}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
	}
	fmt.Println(link, "is up")
}
