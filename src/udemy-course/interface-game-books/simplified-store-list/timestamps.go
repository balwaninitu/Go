package main

import (
	"strconv"
	"time"
)

type timestamp struct {
	time.Time
}

func (ts timestamp) string() string {
	if ts.IsZero() {
		return "unknown"
	}

	const layout = "2006/01"
	return ts.Format(layout)
}

func toTimestamp(v interface{}) (ts timestamp) {
	var t int

	switch v := v.(type) {
	case int:
		t = v
		//fmt.Println("int  -> ")
	case string:
		//fmt.Println("string  -> ")

		t, _ = strconv.Atoi(v)

	}

	ts.Time = time.Unix(int64(t), 0)
	return ts
}
