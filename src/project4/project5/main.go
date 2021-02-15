package main

import (
	"fmt"
	"time"
)

func main() {

	input := "2018-04-24"
	layout := "2006-01-02"
	t, _ := time.Parse(layout, input)
	fmt.Println(t)                       // 2018-04-24 00:00:00 +0000 UTC
	fmt.Println(t.Format("02-Jan-2006")) // 24-Apr-2018

	inputTime := "16:00"
	layoutTime := "15:04:05"
	tm, _ := time.Parse(layoutTime, inputTime)
	fmt.Println(tm)                      // 2018-04-24 00:00:00 +0000 UTC
	fmt.Println(t.Format("02-Jan-2006")) // 24-Apr-2018

}
