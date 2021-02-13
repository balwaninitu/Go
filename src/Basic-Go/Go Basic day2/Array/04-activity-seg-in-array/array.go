package main

import "fmt"

func main() {

	var operatinSystemList [6]string
	operatinSystemList[0] = "WindowsXP"
	operatinSystemList[1] = "Linux 1.0"
	operatinSystemList[2] = "Raspbian Teensy"
	operatinSystemList[3] = "MacOS"
	operatinSystemList[4] = "IOS"
	operatinSystemList[5] = "Google Android"

	for i, v := range operatinSystemList {
		fmt.Printf("Len of operating system %s is %d\n", v, len(operatinSystemList[i]))
	}

	operatinSystemList[0] = "window10"
	operatinSystemList[1] = "Linux 16.04"
	operatinSystemList[2] = "Raspbian Teensy"

	fmt.Println(operatinSystemList)

	var NewOperatingSystemList [9]string

	for i := 0; i < len(operatinSystemList); i++ {
		operatinSystemList[i] = NewOperatingSystemList[i]
	}

	NewOperatingSystemList[6] = "Ubuntu"
	NewOperatingSystemList[7] = "MS-Dos"
	NewOperatingSystemList[8] = "Solaris"
	fmt.Println(NewOperatingSystemList)
	fmt.Println(NewOperatingSystemList[0:3])
	fmt.Println(NewOperatingSystemList[3:6])
	fmt.Println(NewOperatingSystemList[6:])
}
