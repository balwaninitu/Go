package main

import (
	"fmt"
	"strconv"
)

func main() {

	type currency struct {
		currencyName string

		conversionRate float64
	}

	m := make(map[string]currency)

	m["USD"] = currency{"US dollar", 1.1318}
	m["JPY"] = currency{"Japanese yen", 121.05}
	m["GBP"] = currency{"Pound sterling", 0.90630}
	m["CNY"] = currency{"Chines yuan renminbi", 7.9944}
	m["SGD"] = currency{"Singapore dollar", 1.5743}
	m["MYR"] = currency{"Malaysia ringgit", 4.8390}

	var currencyFrom string
	var currencyTo string
	var currencyAmount string

	fmt.Println("enter curreny to convert from")
	fmt.Scanln(&currencyFrom)

	fmt.Println("enter curreny to convert to")
	fmt.Scanln(&currencyTo)

	fmt.Println("enter curreny amount")
	fmt.Scanln(&currencyAmount)

	currencyAmountVal, _ := strconv.ParseInt(currencyAmount, 10, 0)

	result := (float64(currencyAmountVal) / m[currencyFrom].conversionRate) * m[currencyTo].conversionRate

	fmt.Printf("%s: %.2f %s", currencyFrom, result, currencyTo)

}
