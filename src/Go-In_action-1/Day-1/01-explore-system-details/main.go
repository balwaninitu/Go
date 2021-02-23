package main

import (
	"fmt"
	"net"
)

func main() {

	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}
	for j, i := range interfaces {
		fmt.Printf("key : %v Interfaces Name: %v\n", j, i.Name)
		fmt.Printf("Interfaces HardwareAddress: %v\n", i.HardwareAddr)

		byName, err := net.InterfaceByName(i.Name)
		if err != nil {
			fmt.Println(err)
		}

		addressess, err := byName.Addrs()
		if err != nil {
			fmt.Println(err)
		}

		for k, v := range addressess {
			fmt.Printf("Interface Address # %vK: %q\n", k, v.String())
		}
	}
}
