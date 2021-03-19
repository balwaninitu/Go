package main

import (
	"errors"
	"fmt"
)

//node is

type Node struct {
	item string
	next *Node
}

type linkedList struct {
	head *Node
	size int
}

func (p *linkedList) get(index int) (string, error) {
	if p.head == nil {
		return "", errors.New("empty Linked list!")
	}
	if index > 0 && index <= p.size {
		currentNode := p.head
		for i := 1; i <= index-1; i++ {
			currentNode = currentNode.next
		}
		item := currentNode.item
		return item, nil

	}
	return "", errors.New("Invalid Index")
}

func (p *linkedList) addNode(name string) error {
	newNode := &Node{
		item: name,
		next: nil,
	}
	if p.head == nil {
		p.head = newNode
	} else {
		currentNode := p.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
	p.size++
	return nil
}

func (p *linkedList) addAtPos(index int, name string) error {
	newNode := &Node{
		item: name,
		next: nil,
	}

	if index > 0 && index <= p.size+1 {
		if index == 1 {
			newNode.next = p.head
			p.head = newNode

		} else {

			currentNode := p.head
			var prevNode *Node
			for i := 1; i <= index-1; i++ {
				prevNode = currentNode
				currentNode = currentNode.next
			}
			newNode.next = currentNode
			prevNode.next = newNode

		}
		p.size++
		return nil
	} else {
		return errors.New("Invalid Index")
	}
}

func (p *linkedList) remove(index int) (string, error) {
	var item string

	if p.head == nil {
		return "", errors.New("empty Linked list!")
	}
	if index > 0 && index <= p.size {
		if index == 1 {
			item = p.head.item
			p.head = p.head.next
		} else {
			var currentNode *Node = p.head
			var prevNode *Node
			for i := 1; i <= index-1; i++ {
				prevNode = currentNode
				currentNode = currentNode.next

			}
			item = currentNode.item
			prevNode.next = currentNode.next
		}
	}
	p.size--
	return item, nil
}

func (p *linkedList) printAllNodes() error {
	currentNode := p.head
	if currentNode == nil {
		fmt.Println("Linked list is empty.")
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
	fmt.Println("Created linked list")
	fmt.Println()

	fmt.Print("Adding nodes to the linked list...\n\n")
	myList.addNode("Mary")
	myList.addNode("Jaina")
	myList.addNode("Xander")
	myList.addNode("Marc")
	fmt.Println("Showing all nodes in the linked list...")
	myList.printAllNodes()
	fmt.Printf("There are %+v elements in the list in totoal.\n", myList.size)
	fmt.Println()

	fmt.Println("Demoing get...")
	item, error := myList.get(1)

	if error == nil {

		fmt.Println(item)
	} else {
		fmt.Println("Invalid Index")
	}
	fmt.Println()
	fmt.Println("Adding at index...")

	myList.addAtPos(1, "Chris")
	myList.printAllNodes()

}
