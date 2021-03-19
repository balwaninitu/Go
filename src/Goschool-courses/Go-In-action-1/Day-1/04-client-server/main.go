package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	StartTCPServer()
	StartTCPClient()

}

//StartTCPServer for starting TCP server method
func StartTCPServer() {
	myListener, err := net.Listen("tcp", ":5331")
	if err != nil {
		log.Fatalln(err)
	}
	defer myListener.Close()

	for {
		conn, err := myListener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}

}

func handle(conn net.Conn) {

	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Received :", string(data))
		retMsg := string(data) + "\n"
		conn.Write([]byte(retMsg))
	}
}

//StartTCPClient  starts a tcp client and waits for user input
func StartTCPClient() {

	conn, err := net.Dial("tcp", "localhost:5331")
	if err != nil {
		log.Fatalln("connection fails", err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("key in your message: ")
		message, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, message+"\n")

		recMessage, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Received ", recMessage)
	}
}
