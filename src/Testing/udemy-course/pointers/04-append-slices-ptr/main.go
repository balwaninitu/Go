package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Slices")

	slices()

}

func slices() {

	dirs := []string{"up", "down", "left", "right"}

	up(dirs)
	fmt.Println(dirs)

	upPtr(&dirs)
	fmt.Println(dirs)

}

func up(list []string) {

	for i := range list {
		list[i] = strings.ToUpper(list[i])
	}

	list = append(list, strings.ToUpper("diagonal"))
}

func upPtr(list *[]string) {

	lv := *list

	for i := range lv {
		lv[i] = strings.ToUpper(lv[i])
	}

	*list = append(*list, strings.ToUpper("diagonal"))
}
