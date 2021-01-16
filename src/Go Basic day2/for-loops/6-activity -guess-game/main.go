package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Name: ")

	text, _ := reader.ReadString('\n')

	fmt.Println(text)

	rand.Seed(time.Now().UnixNano())

	//	x := rand.Intn(100 + 1)

}
