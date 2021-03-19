package main

import (
	"errors"
	"fmt"
)

//Node is
type Node struct {
	item string
	prev *Node
	next *Node
}

type dblinkedList struct {
	head *Node
	tail *Node
	size int
}

func (l *dblinkedList) addNode(name string) error {
	newNode := &Node{
		item: name,
		prev: nil,
		next: nil,
	}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		currrentNode := l.head
		for currrentNode.next != nil {
			currrentNode = currrentNode.next
		}
		newNode.prev = currrentNode
		currrentNode.next = newNode
		l.tail = newNode
	}
	l.size++
	return nil
}

func (l *dblinkedList) printAllNode() error {
	currrentNode := l.head
	if currrentNode == nil {
		fmt.Println("DB list is empty")
		return nil
	}
	fmt.Printf("%+v\n", currrentNode.item)
	for currrentNode.next != nil {
		currrentNode = currrentNode.next
		fmt.Printf("%+v\n", currrentNode.item)
	}
	return nil
}

func (l *dblinkedList) printAllNodeReverse() error {
	currrentNode := l.tail
	if currrentNode == nil {
		fmt.Println("db list is empyt")
		return nil
	}
	fmt.Printf("%+v\n", currrentNode.item)
	for currrentNode.prev != nil {
		fmt.Printf("%+v\n", currrentNode.prev.item)
		currrentNode = currrentNode.prev
	}
	return nil
}

func (l *dblinkedList) addAtPos(index int, name string) error {
	newNode := &Node{
		item: name,
		next: nil,
		prev: nil,
	}
	if index == 1 {
		l.head.prev = newNode
		newNode.next = l.head
		l.head = newNode
		return nil

	} else if index > 0 && index <= l.size+1 {
		currentNode := l.head
		var prevNode *Node
		for i := 1; i <= index-1; i++ {
			prevNode = currentNode
			currentNode = currentNode.next
		}
		currentNode.prev = currentNode
		newNode.next = currentNode
		prevNode.next = newNode
		newNode.prev = prevNode

	}
	return errors.New("invalid Index")
}

func (l *dblinkedList) removeHead() error {
	currentNode := l.head
	if currentNode == nil {
		fmt.Println("list is empty")
		return nil
	}
	if currentNode.next != nil {
		l.head = currentNode.next
		l.head.prev = nil
	}
	l.size--
	return nil

}

func (l *dblinkedList) removeTail() error {
	currentNode := l.tail
	if currentNode == nil {
		fmt.Println("list is empty")
		return nil
	}
	if currentNode.prev != nil {
		l.tail = currentNode.prev
		l.tail.next = nil
	}
	l.size--
	return nil
}

func (l *dblinkedList) removeAtIndex(index int) error {
	currentNode := l.head
	if currentNode == nil {
		return errors.New("db is empty")
	}
	if index > 0 && index <= l.size {
		if index == 1 {
			l.removeHead()
		} else {
			var currentNode *Node = l.head
			var prevNode *Node
			for i := 1; i <= index-1; i++ {
				prevNode = currentNode
				currentNode = currentNode.next
			}
			prevNode.next = currentNode.next
			currentNode.next.prev = prevNode
		}

	}
	l.size--
	return nil
}

func main() {

	mydblinkedList := dblinkedList{nil, nil, 0}

	mydblinkedList.addNode("hi")
	mydblinkedList.addNode("hello")
	mydblinkedList.addNode("benjour")
	mydblinkedList.addNode("namaste")
	mydblinkedList.addNode("hola")
	//mydblinkedList.addAtPos(4, "ni hao")
	mydblinkedList.printAllNode()
	mydblinkedList.removeHead()
	fmt.Println("after removal of head")
	mydblinkedList.printAllNode()
	mydblinkedList.removeTail()
	fmt.Println("after tail remove")
	mydblinkedList.printAllNode()
	mydblinkedList.removeAtIndex(2)
	fmt.Println("after remove at index")
	mydblinkedList.printAllNode()

	//mydblinkedList.printAllNodeReverse()

	// 	fmt.Println(mydblinkedList.head.item)
	// 	fmt.Println(mydblinkedList.head.next.item)
	// 	fmt.Println(mydblinkedList.tail.prev.item)
	// 	fmt.Println(mydblinkedList.tail.item)

}
