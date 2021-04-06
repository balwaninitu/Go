package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Result struct {
	Success   bool
	Timestamp int
	Base      string
	Date      string
	Rates     map[string]float64
}
type Error struct {
	Success bool
	Error   struct {
		Code int
		Type string
		Info string
	}
}

func main() {
	url := "http://data.fixer.io/api/latest?access_key=5c53af088d12b938fb8f47eff0b8ab96"

	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var result Result
			json.Unmarshal(body, &result)
			if result.Success {
				for i, v := range result.Rates {
					fmt.Println(i, v)
				}
			} else {
				var err Error
				json.Unmarshal(body, &err)
				fmt.Println(err.Error.Info)
			}
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
	fmt.Println("Done")
}
