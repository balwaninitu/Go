package main

import "fmt"

type BinaryNode struct {
	item  string      // to store the data item
	left  *BinaryNode // pointer to point to left node
	right *BinaryNode // pointer to point to right node
}

type BST struct {
	root *BinaryNode
}

func (bst *BST) insertNode(t **BinaryNode, item string) error {

	if *t == nil {
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
func (bst *BST) insert(item string) {
	bst.insertNode(&bst.root, item)

}

func (bst *BST) inOrderTraverse(t *BinaryNode) {
	if t != nil {
		bst.inOrderTraverse(t.left)
		fmt.Println(t.item)
		bst.inOrderTraverse(t.right)
	}
}

func (bst *BST) inOrder() {
	bst.inOrderTraverse(bst.root)
}

func main() {
	bst := &BST{nil}
	eg := []string{"Andy", "Aiken", "Zander", "Jaina", "Mandy"}
	for _, item := range eg {
		bst.insert(item)
	}
	fmt.Println("InOrder Traversal...")
	bst.inOrder()

}
