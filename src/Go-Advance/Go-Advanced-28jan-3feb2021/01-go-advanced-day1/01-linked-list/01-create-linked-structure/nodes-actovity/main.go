package main

import "fmt"

//Node is
type Node struct {
	item string
	next *Node
}

type linkedList struct {
	head *Node
	size int
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

}
