package main

import "fmt"

//3 tower holding name and stack of rings
type tower struct {
	name string
	ring stack
}

func addRings(t *tower, ring int) {
	for ringSize := ring; ringSize > 0; ringSize-- {
		t.ring.push(ringSize)
	}
}

//stack reperesents rings on source tower
type stack struct {
	rings []int
}

//method to push ring on the rods
func (s *stack) push(ring int) {
	s.rings = append(s.rings, ring)
}

//method to reomve top ring from the stack
func (s *stack) pop() int {
	if len(s.rings) == 0 {
		return 0
	}
	//save last ring
	lastRing := s.rings[len(s.rings)-1]

	s.rings = s.rings[:len(s.rings)-1]
	return lastRing

}

func hanoi(n int, from *tower, to *tower, aux *tower) {
	//if there is one disc it get pop and push to dest
	if n == 1 {
		oneRing := from.ring.pop()

		to.ring.push(oneRing)
		fmt.Printf("Moved ring %d from %v to %v\n", n, from.name, to.name)
		return
	}

	hanoi(n-1, from, aux, to)
	oneRing := from.ring.pop()
	to.ring.push(oneRing)
	fmt.Printf("Moved ring %d from %v to %v\n", n, from.name, to.name)

	hanoi(n-1, aux, to, from)

}

func main() {

	var from tower
	var aux tower
	var to tower

	from.name = "a"
	aux.name = "b"
	to.name = "c"

	addRings(&from, 3)

	//rings on tower in the beginning
	fmt.Println("-------Beginning-----------")
	fmt.Printf("From %v\n", from)
	fmt.Printf("Aux %v\n", aux)
	fmt.Printf("To %v\n", to)

	hanoi(3, &from, &to, &aux)

	//after calling recursive function
	fmt.Println("-------End-----------")
	fmt.Printf("From %v\n", from)
	fmt.Printf("aux %v\n", aux)
	fmt.Printf("to %v\n", to)

}
