// Package binarytree contains Binary Search Tree data structure used for quick search for identification and username information, and all the methods that operate on it.
package btree

import (
	"fmt"
	"sync"
)

// The Binary Node for Binary Search Tree
type BinaryNode struct {
	item  string      // to store the data item
	left  *BinaryNode // pointer to point to left node
	right *BinaryNode // pointer to point to right node
}

// The Binary Search Tree Data Structure with root, and mutex
type BST struct {
	Root *BinaryNode
	mu   sync.Mutex
}

// Function to reset the BST to nil
func (bst *BST) Reset() {
	bst.Root = nil
}

// This is the wrapper function to search for an item in the BST
// It takes in an item of type string
// It return the node which stores the item
func (bst *BST) Search(item string) *BinaryNode {
	return bst.searchNode(bst.Root, item)
}

// This is a recursive function to search for an item in the BST
// It takes in a BST node and an item of type string to be search
// The function recursively parses the tree, until it either finds the item, or reaches the end of the BST
// It returns nil if item is not found
// It returns the binary node which stores the item, if item is found
func (bst *BST) searchNode(t *BinaryNode, item string) *BinaryNode {
	if t == nil {
		return nil
	} else {
		if t.item == item {
			return t
		} else {
			if item < t.item {
				return bst.searchNode(t.left, item)
			} else {
				return bst.searchNode(t.right, item)
			}
		}
	}
}

// This is the wrapper function to insert an item in the BST
// It takes in an item of type string and inserts into the BST
// It uses mutex lock and generates a done message when the operation is complete
func (bst *BST) Insert(item string, wglocal *sync.WaitGroup) {
	defer wglocal.Done()
	defer func() {
		if r := recover(); r != nil {
			println("Panic:" + r.(string))
		}
	}()
	bst.mu.Lock()
	{
		bst.insertNode(&bst.Root, item)
	}
	bst.mu.Unlock()
	return
}

// This is a recursive function to insert an item in the BST
// It takes in a BST node and an item of type string to be inserted
// The function recursively parses the tree, until it reaches the node where item shoudl be inserted
// It returns nil if item is inserted successfully
func (bst *BST) insertNode(t **BinaryNode, item string) error {

	if *t == nil {
		// that is if the value inside the memory pointed to by t which is address of bst.Root
		// &bst.Root is nil, then there are no nodes to the tree
		newNode := &BinaryNode{
			item:  item,
			left:  nil,
			right: nil,
		}
		*t = newNode
		return nil
	}

	if item < (*t).item {
		bst.insertNode(&((*t).left), item)
	} else {
		bst.insertNode(&((*t).right), item)
	}

	return nil
}

// This is a wrapper function to traverse the BST in order
func (bst *BST) InOrder() {
	bst.inOrderTraverse(bst.Root)
}

// This is a recursive function to traverse the BST in order
func (bst *BST) inOrderTraverse(t *BinaryNode) {
	if t != nil {
		bst.inOrderTraverse(t.left)
		fmt.Println(t.item)
		bst.inOrderTraverse(t.right)
	}
}

// This function is a wrapper function to delete an item from the BST tree
// It issues a done to wg waitgroup for running concurrently
// It uses the mutex lock during the deletion process
func (bst *BST) Delete(item string, wglocal *sync.WaitGroup) {
	defer wglocal.Done()
	defer func() {
		if r := recover(); r != nil {
			println("Panic:" + r.(string))
		}
	}()
	bst.mu.Lock()
	{
		bst.deleteNode(bst.Root, item)
	}
	bst.mu.Unlock()
	return
}

// This function is a recursive function to delete an item
// It iterates over the BST tree to find the item to be deleted
func (bst *BST) deleteNode(t *BinaryNode, item string) *BinaryNode {
	if t == nil {
		return nil
	}
	if item < t.item {
		t.left = bst.deleteNode(t.left, item)
		return t
	}
	if item > t.item {
		t.right = bst.deleteNode(t.right, item)
		return t
	}
	if t.left == nil && t.right == nil {
		t = nil
		return nil
	}
	if t.left == nil {
		t = t.right
		return t
	}
	if t.right == nil {
		t = t.left
		return t
	}
	leftmostrightside := t.right
	for {
		//find smallest value on the right side
		if leftmostrightside != nil && leftmostrightside.left != nil {
			leftmostrightside = leftmostrightside.left
		} else {
			break
		}
	}
	t.item = leftmostrightside.item
	t.right = bst.deleteNode(t.right, t.item)
	return t
}
