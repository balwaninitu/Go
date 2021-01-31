package main

import "fmt"

//Node is...
type Node struct {
	name string
	next *Node
}

type linkedList struct {
	head *Node
	size int
}

func main() {
	var m = linkedList{}
	m.addNode("Ina")
	m.addNode("Lina")
	m.addNode("Tina")
	m.addNode("Pina")
	m.printAllNode()

}

func (l *linkedList) addNode(name string) error {
	newNode := &Node{
		name: name,
		next: nil,
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
	if currentNode == nil {
		fmt.Println("Linked list is empty")
		return nil
	}

	fmt.Printf("%+v\n", currentNode.name)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", currentNode.name)
	}
	return nil
}
