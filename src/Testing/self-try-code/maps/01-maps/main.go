package main

import "fmt"

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Li"])

	fmt.Println(elements["O"])

	// v, ok := elements["Ne"]
	// if ok {
	// 	fmt.Println(v)
	// } else {
	// 	fmt.Println("not exist")
	// }

	var name string

	fmt.Println("which element you are looking for")
	fmt.Scanln(&name)

	v, ok := elements[name]
	if ok {
		fmt.Println(v)
	}
	if !ok {
		fmt.Println("not exist")
	}
}
