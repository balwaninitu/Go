package main

import (
	"fmt"
	"strconv"
	"time"
)

type book struct {
	title     string
	price     money
	published interface{}
}

func (b book) print() {
	p := format(b.published)

	fmt.Printf("%-15s: %s - (%v)\n", b.title, b.price.string(), p)
}

func format(v interface{}) string {
	var t int

	switch v := v.(type) {
	case int:
		t = v
		//fmt.Println("int  -> ")
	case string:
		//fmt.Println("string  -> ")

		t, _ = strconv.Atoi(v)
	default:
		//fmt.Println("nil  -> ")
		return "unknown"
	}

	// book{title: "hungry witch", price: 24.5, published: 118281600},

	const layout = "2006, Jan" // can change to "01/2006" or "2006/01"

	u := time.Unix(int64(t), 0)
	return u.Format(layout)
}
