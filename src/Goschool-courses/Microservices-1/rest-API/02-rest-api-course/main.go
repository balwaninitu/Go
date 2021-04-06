package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL = "http://localhost:5000/api/v1/courses"
const key = "2c78afaf-97da-4816-bbee-9ad239abb296"

func getCourse(code string) {
	url := baseURL
	if code != "" {
		url = baseURL + "/" + code + "?key=" + key
	} else {
		url = baseURL + "/" + code + "?key=" + key
	}
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		defer response.Body.Close()

		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))

	}
}

func addCourse(code string, jsonData map[string]string) {
	jsonValue, _ := json.Marshal(jsonData)
	response, err := http.Post(baseURL+"/"+code+"?key="+key,
		"application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		defer response.Body.Close()

		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))

	}
}

func updateCourse(code string, jsonData map[string]string) {
	jsonValue, _ := json.Marshal(jsonData)
	request, err := http.NewRequest(http.MethodPut,
		baseURL+"/"+code+"?key="+key,
		bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		defer response.Body.Close()

		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))

	}
}

func deleteCourse(code string) {
	request, err := http.NewRequest(http.MethodDelete,
		baseURL+"/"+code+"?key="+key, nil)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		defer response.Body.Close()

		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))

	}
}

func main() {
	jsonData := map[string]string{"title": "Applied Go Programming"}
	addCourse("IOT210", jsonData)
	// 201 // 201 - Course added: IOT210
	jsonData = map[string]string{"title": "Applied Python Programming"}
	addCourse("IOT210", jsonData)
	// 201 // 201 - Course added: IOT201
	jsonData = map[string]string{"title": "Go Concurrency Programming"}
	updateCourse("PRG01", jsonData)
	// 202 // 202 – Course updated: IOT210
	getCourse("") // get all courses
	// 200
	// {"IOT201":{"Title":"Applied Python Programming"},
	// "IOT210":{"Title":"Go Concurrency Programming"}}
	getCourse("PRG01") // get a specific course
	//200
	// {"Title":"Go Concurrency Programming"}
	getCourse("PRG02")

	deleteCourse("PRG02")
	// 202 //202 – Course deleted: IOT201
	getCourse("") // get all courses
	// 200
	// {"IOT210":{"Title":"Go Concurrency Programming"}}

	// //courses = make(map[string]courseInfo)
	// router := mux.NewRouter()
	// router.HandleFunc("/baseURL", getCourse(code))
	// router.HandleFunc("/api/v1/courses", allcourses)
	// router.HandleFunc("/baseURL/{courseid}", getCourse()).GET(
	// 	"GET")
	// fmt.Println("Listening at port 5000")
	// log.Fatal(http.ListenAndServe(":5000", router))
}
