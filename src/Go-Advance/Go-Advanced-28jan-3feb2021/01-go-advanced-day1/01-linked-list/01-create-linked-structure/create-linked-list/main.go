package main

import "fmt"

//Node is
type Node struct {
	item int
	next *Node
}

type linkedList struct {
	head *Node
	size int
}

func (l *linkedList) addNode(item int) error {
	newNode := &Node{
		item: item,
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
	if l.head == nil {
		fmt.Println("Linkedlist is emty")
		return nil
	}
	fmt.Printf("%+v\n", currentNode.item)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", currentNode.item)
	}
	return nil

}

func main() {

	myList := &linkedList{nil, 0}

	myList.addNode(10)
	myList.addNode(20)
	myList.addNode(30)
	myList.addNode(40)
	fmt.Println("Showing all nodes in the linked list...")
	myList.printAllNode()
	fmt.Printf("There are total %d items in the list\n", myList.size)
}
