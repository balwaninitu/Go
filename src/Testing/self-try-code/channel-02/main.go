package main

import "fmt"

func chars(channel chan string, s []string) {
	for _, i := range s {
		channel <- i

	}
	close(channel)

}

func main() {

	ch := make(chan string)

	s := []string{"Riddham", "kaneesha", "sadhna", "Gauri"}

	go chars(ch, s)

	for {

		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println(v)

	}

}
