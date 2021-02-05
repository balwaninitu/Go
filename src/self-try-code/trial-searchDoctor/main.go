package main

import "fmt"

type details struct {
	availableDays []string
	availableTime []string
}

func (d *details) searchDoctor(s []string, n int, target string) {
	for _, v := range s {
		if v == target {

		}
	}

}

func search(data []map[string]details, target string) {

}

func searches(mp map[string]details, target string) map[string]details {
	for _, val := range mp {
		for _, v := range val.availableDays {
			if v == target {

			}
			// fmt.Printf("Value :%s is available\n", v)
			// fmt.Printf("Index :%d is available\n", i)

		}
		//fmt.Printf("%s is available\n", key)

		return mp

	}
	return nil
}

var k string
var v details

func main() {

	detail1 := details{availableDays: []string{"Mon", "Tue", "Fri"}, availableTime: []string{"1pm", "2pm", "4pm"}}
	detail2 := details{availableDays: []string{"Mon", "Wed", "Thurs"}, availableTime: []string{"1pm", "2pm", "3pm"}}
	detail3 := details{availableDays: []string{"Mon", "Tue", "Wed", "Fri"}, availableTime: []string{"3pm", "4pm", "5pm"}}

	fmt.Println(detail2.availableTime)
	fmt.Println(detail3)

	m := make(map[string]details)
	m["Dr1"] = detail1
	m["Dr2"] = detail2
	m["Dr3"] = detail3
	for k, v = range m {
		fmt.Printf("Doctor:%s Avalable on:%s time:%s\n", k, v.availableDays, v.availableTime)
	}

	// //customised search/lookup of Doctor based on time and date
	var custName string
	//var custDoctorName string
	var custDay string
	var custTime string
	fmt.Println("Enter your name")
	fmt.Scanln(&custName)
	fmt.Println("Pleae select day of appointmen")
	fmt.Scanln(&custDay)
	fmt.Println("Pleae select time of appointmen")
	fmt.Scanln(&custTime)

	// str := searches(m, custDay)

	// fmt.Println(str)

	// ptr := &details{availableDays: []string{"Mon", "Tue", "Fri"}, availableTime: []string{"1pm", "2pm", "4pm"}}
	// ptr.searchDoctor(v.availableDays, len(v.availableDays), custDay)

	// fmt.Printf("%s is available on day %s\n", m[custDay], custDay)
	// value := "mon"
	// var check bool
	// _, check = m[v.availableDays["mon"]]
	// fmt.Println("present", value)
	// fmt.Println("present", check)

}
