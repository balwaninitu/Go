package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sidelength float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height

}

func (s square) getArea() float64 {
	return s.sidelength * s.sidelength

}

func printArea(s shape) {
	fmt.Println(s.getArea())

}

func main() {

	t := triangle{base: 10, height: 10}
	s := square{10}

	printArea(t)
	printArea(s)

}
