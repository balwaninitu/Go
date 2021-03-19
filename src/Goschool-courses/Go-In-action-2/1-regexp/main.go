package main

import (
	"fmt"
	"regexp"
)

func main() {

	match, _ := regexp.MatchString("p[a-z]+ch", "peach")
	fmt.Println(match)

}
