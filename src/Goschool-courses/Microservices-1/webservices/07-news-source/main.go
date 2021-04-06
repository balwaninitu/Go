package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func fetchData(url string) {
	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var result map[string]interface{}

			json.Unmarshal([]byte(body), &result)

			//articles is now of type []interface{}
			articles := result["articles"]

			for _, article := range articles.([]interface{}) {
				//get the source
				source := article.(map[string]interface{})["source"]
				fmt.Println("=================")
				//get the name
				fmt.Println(source.(map[string]interface{})["name"])
				fmt.Println("=================")

				//get the title
				title := article.(map[string]interface{})["title"]
				fmt.Println(title)

			}
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}

}

func main() {

	fetchData("https://newsapi.org/v2/top-headlines?" +
		"country=us&category=business&apiKey=10a589fc2f0541e7845887e4197ed8fb")

	fmt.Println("Done!")
	fmt.Scanln()
}
