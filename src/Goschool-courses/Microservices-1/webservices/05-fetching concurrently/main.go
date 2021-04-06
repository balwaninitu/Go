package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var (
	apis map[int]string
	wg   sync.WaitGroup
)

func fetchData(API int) {
	defer wg.Done()
	url := apis[API]
	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			switch API {
			case 1:
				if result["success"] == true {
					fmt.Println(result["rates"].(map[string]interface{})["USD"])
				} else {
					fmt.Println(result["error"].(map[string]interface{})["info"])
				}
			case 2:
				if result["main"] != nil {
					fmt.Println(result["main"].(map[string]interface{})["temp"])
				} else {
					fmt.Println(result["message"])
				}
			}
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}

func main() {
	apis = make(map[int]string)
	apis[1] = "http://data.fixer.io/api/latest?access_key=5c53af088d12b938fb8f47eff0b8ab96"
	apis[2] = "http://api.openweathermap.org/data/2.5/weather?q=SINGAPORE&appid=a32565ee3f9c717d0439e3c32b1b5074"

	wg.Add(2)
	go fetchData(1)
	go fetchData(2)

	wg.Wait()

	//fmt.Scanln()
}
