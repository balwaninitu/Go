package main

import (
	"fmt"
	"strconv"
)

func main() {

	local := 10

	local = incrN(local)
	showN(local)

	local = incrBy(local, 5)
	showN(local)

	_, err := incrBystr(local, "two")
	if err != nil {

		fmt.Printf("Error %s\n", err)
	}

	local, _ = incrBystr(local, "2")
	showN(local)

	showN(incrBy(local, 2))

	local = sanitize(incrBystr(local, "2"))
	showN(local)

	local = limit(incrBy(local, 5), 2000)
	showN(local)
}

func showN(n int) {

	fmt.Println("show N", n)
}

func incrN(n int) int {

	n++
	return n
}

func incrBy(n, factor int) int {

	return n * factor
}

func incrBystr(n int, factor string) (int, error) {

	m, err := strconv.Atoi(factor)
	n = incrBy(n, m)

	return n, err
}

func sanitize(n int, err error) int {

	if err != nil {
		return 0
	}
	return n
}

func limit(n, lim int) (m int) {

	m = n
	if m >= lim {

		m = lim
	}

	return
}
