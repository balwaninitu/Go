package main

import (
	"fmt"
	"math"
)

type generic interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	length float64
	width  float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.length * r.width
}

func calculate(g generic) {
	fmt.Println(g)
	fmt.Println(g.area())

}

func main() {

	c := circle{radius: 5}

	r := rectangle{length: 4, width: 3}

	calculate(c)
	calculate(r)

}
