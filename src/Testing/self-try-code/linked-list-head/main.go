package main

import (
	"errors"
	"fmt"
)

//Node is...
type Node struct {
	item string
	next *Node
}

type linkedList struct {
	head *Node
	size int
}

func (l *linkedList) addNode(item string) error {
	newNode := &Node{
		item,
		nil,
	}
	if l.head == nil {
		l.head = newNode
	} else {
		currentNode := l.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
	l.size++
	return nil
}

func (l *linkedList) printAllNode() error {
	currentNode := l.head
	if l.head == nil {
		fmt.Println("linked list is empty")
		return nil
	}

	fmt.Printf("%+v\n", currentNode.item)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", currentNode.item)
	}
	return nil
}

func (l *linkedList) getIndex(index int) (string, error) {
	if l.head == nil {
		return "", errors.New("empty list")
	}
	if index > 0 && index <= l.size {
		currentNode := l.head
		for i := 1; i <= (index - 1); i++ {
			currentNode = currentNode.next

		}
		item := currentNode.item
		return item, nil
	}
	return "", errors.New("invalid index")

}

func main() {

	mylinkedlist := linkedList{}

	mylinkedlist.addNode("hi")
	// mylinkedlist.addNode("hello")
	// mylinkedlist.addNode("hola")
	// mylinkedlist.printAllNode()

	item, error := mylinkedlist.getIndex(1)
	if error == nil {
		fmt.Println(item)
	} else {
		fmt.Println("invalid index")
	}
	mylinkedlist.getIndex(2)

}
