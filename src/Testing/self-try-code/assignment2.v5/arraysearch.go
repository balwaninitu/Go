package main

import (
	"errors"
	"fmt"
)

func search(doctors []details, n int, target string) error {
	for _, v := range doctors {
		if v.name == target {

			fmt.Printf("Doctor Name: %s is available on %s", v.name, v.availableDay)
			return nil
		}
	}
	fmt.Printf("%s is not available", target)

	return errors.New("Invalid name")

}
