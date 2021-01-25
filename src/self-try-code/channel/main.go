package main

import "fmt"

func looping(channel chan int) {
	for i := 0; i < 10; i++ {
		channel <- i
	}
	close(channel)
}

func main() {

	ch := make(chan int)

	go looping(ch)

	for {

		v, ok := <-ch

		if ok == false {
			break
		}

		fmt.Println("Received", ok, v)
	}

}
