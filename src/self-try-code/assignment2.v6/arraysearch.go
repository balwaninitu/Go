package main

import (
	"fmt"
	"time"
)

func search(doctors []doctorDetails, n int, doctorName string) error {
	for i, d := range doctors {
		if d.name == doctorName {
			i = i + 1
			fmt.Printf("%s is available on %v\n", doctorName, d.availableTime.Format(time.ANSIC))
			return nil
		}

	}
	fmt.Printf("%s is not available\nYou can check doctor names from list\n", doctorName)
	return nil

}
