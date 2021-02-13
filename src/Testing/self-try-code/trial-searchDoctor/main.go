package main

import "fmt"

type details struct {
	availableDays []string
	availableTime []string
}

func search(s []string, n int, target string) bool {
	for _, val := range s {
		if val == target {
			return true
		}
	}
	return false
}

var k string
var v details
var detail1 details
var detail2 details
var detail3 details

func main() {

	detail1 = details{availableDays: []string{"Mon", "Tue", "Fri"}, availableTime: []string{"1pm", "2pm", "4pm"}}
	detail2 = details{availableDays: []string{"Mon", "Wed", "Thurs"}, availableTime: []string{"1pm", "2pm", "3pm"}}
	detail3 = details{availableDays: []string{"Mon", "Tue", "Wed", "Fri"}, availableTime: []string{"3pm", "4pm", "5pm"}}

	// fmt.Println(detail2.availableTime)
	// fmt.Println(detail3)

	m := make(map[string]details)
	m["Dr1"] = detail1
	m["Dr2"] = detail2
	m["Dr3"] = detail3
	for k, v = range m {
		fmt.Printf("Doctor:%s Avalable on:%s time:%s\n", k, v.availableDays, v.availableTime)
	}

	var target1 string
	var target2 string
	for _, dVal := range detail1.availableDays {
		if dVal == target1 {
			fmt.Printf("Dr1 is avaialble on %s\n", target1)
		}
	}
	for _, dVal := range detail2.availableDays {
		if dVal == target1 {
			fmt.Printf("Dr2 is avaialble on %s\n", target1)
		}
	}
	for _, dVal := range detail3.availableDays {
		if dVal == target1 {
			fmt.Printf("Dr3 is avaialble on %s\n", target1)
		}
	}
	for _, dVal := range detail1.availableTime {
		if dVal == target2 {
			fmt.Printf("Dr1 is avaialble on %s\n", target2)
		}
	}
	for _, dVal := range detail2.availableTime {
		if dVal == target1 {
			fmt.Printf("Dr2 is avaialble on %s\n", target1)
		}
	}
	for _, dVal := range detail3.availableDays {
		if dVal == target1 {
			fmt.Printf("Dr3 is avaialble on %s\n", target1)

		}

	}

}
