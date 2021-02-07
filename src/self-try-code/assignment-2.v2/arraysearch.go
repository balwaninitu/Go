package main

import "fmt"

func appendSlice(doctors []doctor) []int {
	var s []int
	var i int
	for i = range doctors {
		s = append(s, i)

	}
	return s
}

func search(s []int, n int, target int) error {
	var i int
	for i = range s {
		if i == target {
			i = i + 1
			fmt.Println("id found")

		}
	}
	return nil

}
