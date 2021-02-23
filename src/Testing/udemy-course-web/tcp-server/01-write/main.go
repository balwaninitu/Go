package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		io.WriteString(conn, "What are you doing there?")
		fmt.Fprintln(conn, "coding?")
		fmt.Fprintf(conn, "%v", "Great!!")

		conn.Close()

	}
}
