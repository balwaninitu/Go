package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

const baseURL = "http://localhost:8080/courses"
const key = "2c78afaf-97da-4816-bbee-9ad239abb296"

var (
	Client *sql.DB
	err    error
)

func init() {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		"user", "password", "127.0.0.1:3306", "users_db")

	Client, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err.Error())
	}
	//defer Client.Close()
	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println(" Connected to Database ")

}

func getCourse(courseId int64) {
	url := baseURL
	//  if code != "" {
	//  	url = baseURL + "/" + code + "?key=" + key
	// } else {
	url = baseURL + "/" + strconv.FormatInt(courseId, 10) + "?key=" + key
	// }
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

func addCourse(jsonData map[string]string) {
	jsonValue, _ := json.Marshal(jsonData)
	response, err := http.Post(baseURL+"/"+"?key="+key,
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

func updateCourse(courseId int64, jsonData map[string]string) {
	jsonValue, _ := json.Marshal(jsonData)
	request, err := http.NewRequest(http.MethodPut,
		baseURL+"/"+strconv.FormatInt(courseId, 10)+"?key="+key,
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

func deleteCourse(courseId int64) {
	request, err := http.NewRequest(http.MethodDelete,
		baseURL+"/"+strconv.FormatInt(courseId, 10)+"?key="+key, nil)
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

	getCourse(125)

	getCourse(126)

	jsonData := map[string]string{"title": "goAction-2"}
	addCourse(jsonData)

	jsonData = map[string]string{"124": "goaction"}
	addCourse(jsonData)

	jsonData = map[string]string{"title": "goarun"}
	updateCourse(138, jsonData)

	deleteCourse(138)

}
