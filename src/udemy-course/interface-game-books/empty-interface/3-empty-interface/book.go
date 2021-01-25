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

	// 	book{title: "petit witch", price: 24.5},

	if v == nil {
		return "unknown"
	}

	var t int

	// book{title: "hungry witch", price: 24.5, published: 118281600},
	if v, ok := v.(int); ok {
		t = v
	}

	// 	book{title: "evil witch", price: 24.5, published: "733622400"},

	if v, ok := v.(string); ok {
		t, _ = strconv.Atoi(v)
	}

	u := time.Unix(int64(t), 0)
	return u.String()
}
