package main

import (
	"errors"
	"fmt"
)

//Node is
type Node struct {
	item string
	next *Node
}

type stack struct {
	top  *Node
	size int
}

func (s *stack) push(name string) error {
	newNode := &Node{
		item: name,
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
	cureentNode := s.top
	if cureentNode == nil {
		fmt.Println("stack is empty")
		return nil
	}
	fmt.Printf("%+v\n", cureentNode.item)
	for cureentNode.next != nil {
		cureentNode = cureentNode.next
		fmt.Printf("%+v\n", cureentNode.item)
	}
	return nil
}

func (s *stack) pop() (string, error) {
	var item string
	if s.top == nil {
		return "", errors.New("stack is empty")
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

func (s *stack) getSize() int {
	return s.size
}

func (s *stack) isEmpty() bool {
	return s.size == 0
}

func main() {

	myStack := stack{nil, 0}
	myStack.push("hi")
	myStack.push("namaste")
	myStack.push("hello")
	myStack.push("hola")
	myStack.printAllNode()
	fmt.Println("Size of stack", myStack.size)
	myStack.getSize()
	// myStack.pop()
	// fmt.Println("after one pop")
	// myStack.printAllNode()
	// myStack.pop()
	// fmt.Println("after second pop")
	// myStack.printAllNode()
	// myStack.pop()
	// fmt.Println("after third pop")
	// myStack.printAllNode()
	// myStack.pop()
	// fmt.Println("after fourth pop")
	// myStack.printAllNode()
	// myStack.pop()
	// fmt.Println("after fifth pop")
	// myStack.printAllNode()

	tempStack := &stack{nil, 0}
	for myStack.isEmpty() == false {
		item, _ := myStack.pop()
		tempStack.push(item)
	}

	fmt.Println("tempstack size:", tempStack.size)
	fmt.Println("mystack size:", myStack.size)

	for tempStack.isEmpty() == false {
		item, _ := tempStack.pop()
		myStack.push(item)
	}
	fmt.Println("tempstack size:", tempStack.size)
	fmt.Println("mystack size:", myStack.size)

}
