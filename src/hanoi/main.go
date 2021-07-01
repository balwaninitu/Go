package main

import (
	"errors"
	"fmt"
)

// counter for number of moves
var cnt int

// Stack struct wraps a go slice
type stack struct {
	a []int
}

// Push takes an element and adds it to stack
func (s *stack) push(e int) {
	// Internally, we append the element to the underlying slice
	s.a = append(s.a, e)
}

// Pop throws the top most element out of stack
// If the stack is empty, it throws an error and returns -1
func (s *stack) pop() (int, error) {
	// Check if stack is empty
	// We can't pop from an empty stack; throw error
	if len(s.a) == 0 {
		return -1, errors.New("empty Stack")
	}

	// Top most element is the last element in underlying slice
	// which lies at index len(slice)-1
	// store this value and delete the element from the slice
	// Note there is no delete method and we need to use append
	// Basically append all elements except the last one
	res := s.a[len(s.a)-1]
	s.a = s.a[:len(s.a)-1]

	return res, nil
}

// Tower type holds the name of the tower and the stack of disks
type tower struct {
	name  string
	disks stack
}

// populateSource will populate the tower with a number of disks
func populateSource(t *tower, n int) {
	// populate the source tower with 4 disks
	// size indicates the size of disk
	// bottom being the largest (size 4) and top the smallest (size 1)
	for size := n; size > 0; size-- {
		t.disks.push(size)
	}
}

// TowerOfHanoi takes the number of disks, and pointer to three towers
func TowerOfHanoi(n int, source *tower, dest *tower, aux *tower) {

	// If only 1 disk, pop it from source, push to destination, and print the move
	// Note that the actual source and destination are changing with recursive calls
	// We are simply panicking on errors for now
	if n == 1 {
		r, err := source.disks.pop()
		if err != nil {
			panic(err)
		}
		dest.disks.push(r)
		cnt++
		fmt.Printf("Moved disk %d from %v to %v\n", n, source.name, dest.name)
		return
	}

	// Moving top n-1 disks from A to B using C as auxiliary
	TowerOfHanoi(n-1, source, aux, dest)

	// Pop the disk from source and push to destination
	// Print this move
	r, err := source.disks.pop()
	if err != nil {
		panic(err)
	}
	dest.disks.push(r)
	cnt++
	fmt.Printf("Moved disk %d from %v to %v\n", n, source.name, dest.name)

	// Move n-1 disks from B to C using A as auxiliary
	TowerOfHanoi(n-1, aux, dest, source)
}

func main() {

	// source, auxiliary and destination towers also called pegs
	var srcPeg tower
	var auxPeg tower
	var destPeg tower

	// set in the name of each peg
	srcPeg.name = "a"
	auxPeg.name = "b"
	destPeg.name = "c"

	// populate the source tower with 4 disks
	populateSource(&srcPeg, 3)

	// Print the initial state of pegs
	fmt.Printf("Source Peg %+v\n", srcPeg)
	fmt.Printf("Auxiliary Peg %+v\n", auxPeg)
	fmt.Printf("Destination Peg %+v\n", destPeg)

	// call TowerOfHanoi with 4 disks and 3 pegs
	TowerOfHanoi(3, &srcPeg, &destPeg, &auxPeg)

	// Print the final state of pegs
	fmt.Printf("Source Peg %+v\n", srcPeg)
	fmt.Printf("Auxiliary Peg %+v\n", auxPeg)
	fmt.Printf("Destination Peg %+v\n", destPeg)

	// Number of moves
	fmt.Printf("Number of moves: %d\n", cnt)
}
