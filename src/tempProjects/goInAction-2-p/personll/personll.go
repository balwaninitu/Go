// Package personll contains the linked list data structure for storing data for each person and all the methods that manipulate it.
package personll

import (
	"errors"
	"fmt"
	"goInAction-2-p/person"
	"sync"
	//user "vaccineappt/source/person"
)

// Linked list node containing an item of type person, and a pointer to next node
type Node struct {
	Item person.Person
	Next *Node
}

// Linked list structure containing the head node, size of linked list, and mutex
type LinkedList struct {
	Head *Node
	Size int
	mu   sync.Mutex
}

// This is a method for linked list struct
// It is used to get a person at the specified index
// It returns person struct and any errors
func (p *LinkedList) Get(index int) (person.Person, error) {
	emptyItem := person.Person{}
	if p.Head == nil {
		return emptyItem, errors.New("Empty Linked list!")
	}
	if index > 0 && index <= p.Size {
		currentNode := p.Head
		for i := 1; i <= index-1; i++ {
			currentNode = currentNode.Next
		}
		item := currentNode.Item
		return item, nil

	}
	return emptyItem, errors.New("Invalid Index")
}

// This is a method for linked list struct
// It is used to add a person  to the linked list
// It takes in a name of type person to be added
// The node for the person is added at the end of the linked list
// Waitgroup done is issued for concurrent processing
// Mutex lock is enabled during addition of node
func (p *LinkedList) AddNode(name person.Person, wglocal *sync.WaitGroup) error {
	defer wglocal.Done()

	defer func() {
		if r := recover(); r != nil {
			println("Panic:" + r.(string))
		}
	}()
	p.mu.Lock()
	{
		newNode := &Node{
			Item: name,
			Next: nil,
		}
		if p.Head == nil {
			p.Head = newNode
		} else {
			currentNode := p.Head
			for currentNode.Next != nil {
				currentNode = currentNode.Next
			}
			currentNode.Next = newNode
		}
		p.Size++

	}
	p.mu.Unlock()
	return nil
}

// This is a method for linked list struct
// It is used to add a person to the linked list at a certain index
// It traverses the linked list, and adds name of type person at location index
func (p *LinkedList) AddAtPos(index int, name person.Person) error {
	newNode := &Node{
		Item: name,
		Next: nil,
	}

	if index > 0 && index <= p.Size+1 {
		if index == 1 {
			newNode.Next = p.Head
			p.Head = newNode

		} else {

			currentNode := p.Head
			var prevNode *Node
			for i := 1; i <= index-1; i++ {
				prevNode = currentNode
				currentNode = currentNode.Next
			}
			newNode.Next = currentNode
			prevNode.Next = newNode

		}
		p.Size++
		return nil
	} else {
		return errors.New("Invalid Index")
	}
}

// This is a method for linked list struct
// It is used to remove a person from the linked list at a certain index
// It traverses the linked list, and removesthe person at location index
// It returns the person that is removed, and any errors if present
func (p *LinkedList) Remove(index int) (person.Person, error) {
	var item person.Person
	emptyItem := person.Person{}

	if p.Head == nil {
		return emptyItem, errors.New("Empty Linked list!")
	}
	if index > 0 && index <= p.Size {
		if index == 1 {
			item = p.Head.Item
			p.Head = p.Head.Next
		} else {
			var currentNode *Node = p.Head
			var prevNode *Node
			for i := 1; i <= index-1; i++ {
				prevNode = currentNode
				currentNode = currentNode.Next

			}
			item = currentNode.Item
			prevNode.Next = currentNode.Next
		}
	}
	p.Size--
	return item, nil
}

// This is a method for linked list struct
// It is used to generate message for admin listing all the users in the linked list
// It traverses the linked list, and returns the usernames for each person and any messages
// This is used by the admin template to display list of all users
func (p *LinkedList) PrintAllUsers() (msg []string, usr []string) {
	var message []string
	var users []string

	count := 1
	currentNode := p.Head
	if currentNode == nil {
		message = append(message, fmt.Sprintf("No users found."))
		return message, users
	}
	message = append(message, fmt.Sprintf("\nListing all usernames:"))
	username := currentNode.Item.Username
	users = append(users, fmt.Sprintf("%s", username))

	count++
	for currentNode.Next != nil {
		currentNode = currentNode.Next
		username = currentNode.Item.Username
		users = append(users, fmt.Sprintf("%s", username))
		count++
	}
	return message, users
}

// This is a method for linked list struct
// It is used to search for a username in the linked list
// It traverses the linked list, and returns the person data for the username
// It also returns the index at which the username is found, and any errors if present
func (p *LinkedList) SearchUserName(username string) (person.Person, int, error) {
	emptyItem := person.Person{}
	index := 1
	if p.Head == nil {
		return emptyItem, -1, errors.New("Empty Linked list!")
	}
	currentNode := p.Head
	for i := 1; i <= p.Size; i++ {
		if currentNode.Item.Username != username {
			currentNode = currentNode.Next
			index++
		} else {
			item := currentNode.Item
			return item, index, nil
		}
	}
	return emptyItem, -1, errors.New("Invalid Username")
}

// This is a method for linked list struct
// It is used to write person data at a specified index in the linked list
// It takes in an index of type int, and thisPerson of type person
// The person data at the index specified is overwritten by contents of thisPerson
// It returns any error if present
func (p *LinkedList) WriteAtIndex(index int, thisPerson person.Person) error {

	if p.Head == nil {
		return errors.New("Empty Linked list!")
	}
	if index > 0 && index <= p.Size {
		currentNode := p.Head
		for i := 1; i <= index-1; i++ {
			currentNode = currentNode.Next
		}
		currentNode.Item = thisPerson
		return nil

	}
	return errors.New("Invalid Index")
}

// This is a method for linked list struct
// It is a wrapper used to write person data at a specified index in the linked list
// It takes in an index and a person
// The person at index specified is overwritten
// It call writeAtIndex function to write the person data at index
func (p *LinkedList) WritePersonData(currentPerson person.Person, currentPersonIndex int) error {
	err := p.WriteAtIndex(currentPersonIndex, currentPerson)
	if err != nil {
		return err
	}
	return nil
}
