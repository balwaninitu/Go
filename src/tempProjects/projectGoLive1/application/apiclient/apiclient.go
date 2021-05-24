/*
Go file containing package apiclient for running Seller API
The apiclient package allows the user to:
- Add item
- Update item
- Delete item
- Retrieve item
The package ignores TLS connection security as self-generated certificates are used for this project.
*/
package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	config "projectGoLive/application/config"
)

//Data structure for each item
type ItemsDetails struct {
	Item     string  `json:"Item"`
	Quantity int     `json:"Quantity"`
	Cost     float64 `json:"Cost"`
	Username string  `json:"Username"`
}

// Variable used only within this package
var buyerapikey string
var sellerapikey string

// Base URL used for connecting to seller REST API
const baseURL = "https://" + config.APIPortNum + "/api/v1/"

//For Seller

// Initialization function
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	// Get the SELLER_API_KEY environment variable
	sellerapikey, _ = os.LookupEnv("SELLER_API_KEY")
	buyerapikey, _ = os.LookupEnv("BUYER_API_KEY")
}

// This function sends a request to the REST API to get one or all Items, and then displays the response.
// It ignores TLS security as REST API server uses self generated certicates
// It takes in the name of the Item to search, of type string
// If code is empty, it sends a request to search all Items
// Upon receiving the response from REST API, it displays the status of the request and the Item details.
func getItem(IN, SN string, isSeller bool) ([]ItemsDetails, bool) {
	var Items []ItemsDetails
	url := ""
	if isSeller {
		url = baseURL + "seller"
	} else {
		url = baseURL + "buyer"
	}
	fmt.Println("URL is ", url)
	/*	// Skipping TLS verification as self generated certificate is used
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
	*/
	if IN != "" {
		//url = baseURL + "/" + SN + "/" + IN + "?key=" + apikey
		url = url + "/" + SN + "/" + IN
	} else if SN != "" {
		//url = baseURL + "/" + SN + "?key=" + apikey
		url = url + "/" + SN
	}
	fmt.Println("URL is ", url)

	//response, err := client.Get(url)
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		defer response.Body.Close()
		data, _ := ioutil.ReadAll(response.Body)
		if response.StatusCode == 200 {
			if IN != "" { // get one Item
				var oneItem ItemsDetails
				err := json.Unmarshal(data, &oneItem)
				if err != nil {
					log.Println(err)
				} else {
					Items = append(Items, oneItem)
					fmt.Println("Details of Item are : ")
					fmt.Printf("Item: \"%s\"\n", oneItem.Item)
					fmt.Printf("Quantity: %d\n", oneItem.Quantity)
					fmt.Printf("Cost: %f\n", oneItem.Cost)
					fmt.Printf("Username: \"%s\"\n", oneItem.Username)
					fmt.Println()
					// return one item in Items array, and true for successful get
					return Items, true
				}
			} else { // all Items
				err := json.Unmarshal(data, &Items)
				if err != nil {
					log.Println(err)
				} else {
					fmt.Println("List of all Items : ")
					for i, item := range Items {
						fmt.Printf("------- %d -------\n", i+1)
						fmt.Printf("Item: \"%s\"\n", item.Item)
						fmt.Printf("Quantity: %d\n", item.Quantity)
						fmt.Printf("Cost: %f\n", item.Cost)
						fmt.Printf("Username: \"%s\"\n", item.Username)
					}
					fmt.Println()
					// return all items in Items array, and true for successful get
					return Items, true
				}
			}
		} else if response.StatusCode == 404 {
			fmt.Println("Item not found. Try again")
			fmt.Println()
		} else {
			fmt.Println(response.StatusCode)
			fmt.Println(string(data))
			fmt.Println()
		}
	}
	// This return is for all errors, Items array will be empty, and false is not successful
	return Items, false
}

// This function sends a request to the REST API to add one Item, and then displays the response.
// It ignores TLS security as REST API server uses self generated certicates
// It takes in the name of the Item to add of type string.
// It also takes in the json data to be sent containing details of the Item to add.
// Upon receiving the response from REST API, it displays the status of the request and if Item has been added successfully.
//func addItem(code string, jsonData map[string]string) {
func addItem(IN, SN string, isSeller bool, si ItemsDetails) bool {
	url := ""
	if isSeller {
		url = baseURL + "seller"
	} else {
		//url = baseURL + "buyer"
		return false
	}
	fmt.Println("URL is ", url)

	/*
		// Skipping TLS verification as self generated certificate is used
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}

		response, err := client.Post(baseURL+"/"+code+"?key="+apikey,
			"application/json", bytes.NewBuffer(jsonValue))
	*/

	jsonValue, _ := json.Marshal(si)

	response, err := http.Post(url+"/"+SN+"/"+IN, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		defer response.Body.Close()
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		if response.StatusCode == 201 {
			fmt.Println("Item added successfully.")
			fmt.Println()
			return true
		} else if response.StatusCode == 409 {
			fmt.Println("Item already exists! Try again.")
			fmt.Println()
		} else {
			fmt.Println(response.StatusCode)
			fmt.Println(string(data))
			fmt.Println()
		}
	}
	return false
}

// This function sends a request to the REST API to update one Item, and then displays the response.
// It ignores TLS security as REST API server uses self generated certicates
// It takes in the name of the Item to update of type string.
// It also takes in the json data to be sent containing details of the Item to update.
// Upon receiving the response from REST API, it displays the status of the request and if Item has been updated successfully.
//func updateItem(code string, jsonData map[string]string) {
func updateItem(IN, SN string, isSeller bool, si ItemsDetails) bool {
	url := ""
	if isSeller {
		url = baseURL + "seller"
	} else {
		url = baseURL + "buyer"
	}
	fmt.Println("URL is ", url)

	jsonValue, _ := json.Marshal(si)
	request, err := http.NewRequest(http.MethodPut,
		url+"/"+SN+"/"+IN,
		bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	/*
		request, err := http.NewRequest(http.MethodPut,
			baseURL+"/"+code+"?key="+apikey,
			bytes.NewBuffer(jsonValue))
		request.Header.Set("Content-Type", "application/json")


		// Skipping TLS verification as self generated certificate is used
		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
		response, err := client.Do(request)*/

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		defer response.Body.Close()
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		if response.StatusCode == 201 {
			fmt.Println("Item not in database. Added as a new Item.")
			fmt.Println()
			return true
		} else if response.StatusCode == 202 {
			fmt.Println("Item updated successfully.")
			fmt.Println()
			return true
		} else {
			fmt.Println(response.StatusCode)
			fmt.Println(string(data))
			fmt.Println()
		}
	}
	return false
}

// This function sends a request to the REST API to delete one Item, and then displays the response.
// It ignores TLS security as REST API server uses self generated certicates
// It takes in the name of the Item to be deleted of type string.
// Upon receiving the response from REST API, it displays the status of the request and if Item has been deleted successfully.
func deleteItem(IN, SN string, isSeller bool) bool {
	url := ""
	if isSeller {
		url = baseURL + "seller"
	} else {
		url = baseURL + "buyer"
	}
	fmt.Println("URL is ", url)

	//request, err := http.NewRequest(http.MethodDelete,baseURL+"/"+code+"?key="+apikey, nil)
	request, err := http.NewRequest(http.MethodDelete, url+"/"+SN+"/"+IN, nil)

	/*	// Skipping TLS verification as self generated certificate is used
		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}*/
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		defer response.Body.Close()
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		if response.StatusCode == 202 {
			fmt.Println("Item deleted successfully.")
			fmt.Println()
			return true
		} else if response.StatusCode == 404 {
			fmt.Println("Item not found. Try again")
			fmt.Println()
		} else {
			fmt.Println(response.StatusCode)
			fmt.Println(string(data))
			fmt.Println()
		}
	}
	return false
}
