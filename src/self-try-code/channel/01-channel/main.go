package main

import "fmt"

func add(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {

	s := []int{1, 2, 3, 7, 9}
	c := make(chan int)

	go add(s, c)

	x := <-c

	fmt.Println(x)

}
