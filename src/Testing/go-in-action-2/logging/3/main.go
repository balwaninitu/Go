package main

import (
	"log"
	"os"
)

func main() {

	file, e := os.OpenFile("sample.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if e != nil {
		log.Fatalln("failed")
	}

	log.SetOutput(file)
	log.Println("Welcome")
}
