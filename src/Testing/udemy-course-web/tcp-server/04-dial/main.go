package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	//dial write
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	//fmt.Fprintf(conn, "I dialed you")

	//dial read

	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))
}
