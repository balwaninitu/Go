package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	n := bufio.NewScanner(os.Stdin)

	for {
		n.Scan()

		fmt.Println("Text:", n.Text())
		fmt.Println("Text:", n.Bytes())

		if n.Text() == "hi" {

			break
		}

	}

}
