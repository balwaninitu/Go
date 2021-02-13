package main

import (
	"errors"
	"fmt"
)

//Node is exported
type Node struct {
	item string
	next *Node
}

type stack struct {
	top  *Node
	size int
}

func (s *stack) push(item string) error {
	newNode := &Node{
		item: item,
		next: nil,
	}
	if s.top == nil {
		s.top = newNode
	} else {
		newNode.next = s.top
		s.top = newNode
	}
	s.size++
	return nil
}

func (s *stack) printAllNode() error {
	pushNode := s.top
	if s.top == nil {
		fmt.Println("stack is empty")
		s.top = pushNode
	}
	fmt.Printf("%+v\n", pushNode.item)
	return nil
}

func (s *stack) pop() (string, error) {
	var item string
	if s.top == nil {
		return "", errors.New("Empty stack")
	}
	item = s.top.item

	if s.size == 1 {
		s.top = nil
	} else {
		s.top = s.top.next
	}
	s.size--
	return item, nil
}

func main() {

	myStack := stack{nil, 0}

	myStack.push("Ina")
	myStack.push("Ina")
	myStack.push("Ina")
	myStack.push("Ina")
	myStack.printAllNode()

	//expression := "{}[]()"

}
