package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2021, time.March, 19, 10, 0, 0, 0, time.Local)
	fmt.Printf("Date %s\n", t.Local())

}
