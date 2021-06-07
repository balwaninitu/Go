package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	dummyhead := &ListNode{}

	currentNode := dummyhead

	//Iterate over both list if not nil

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			//next node in head will be 1(l1.val)
			currentNode.Next = l1
			//l1 taken up already hence it become l1.next
			l1 = l1.Next
			//currentnode will become next i.e 2
			currentNode = currentNode.Next

		} else {
			currentNode.Next = l2
			l2 = l2.Next
			currentNode = currentNode.Next
		}
	}

	if l1 != nil {
		currentNode.Next = l1

	} else if l2 != nil {
		currentNode.Next = l2
	}
	fmt.Println(dummyhead.Next)
	return dummyhead.Next
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {

	// Base cases
	if (l1 == nil) && (l2 == nil) {

		// Both l1 and l2 are empty
		return nil

	} else if l1 == nil {

		// Only l1 is empty
		return l2

	} else if l2 == nil {

		// Only l2 is empty
		return l1
	}

	// General cases
	if l1.Val < l2.Val {

		// l1 is smaller than l2
		l1.Next = mergeTwoLists2(l1.Next, l2)
		return l1

	} else {

		// l2 is smaller than l1
		l2.Next = mergeTwoLists2(l1, l2.Next)
		return l2

	}

}

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 != nil {
		return l2
	}
	if l1 != nil && l2 == nil {
		return l1
	}
	if l1 == nil && l2 == nil {
		return nil
	}
	newNode := new(ListNode)
	if l1.Val >= l2.Val {
		newNode.Val = l2.Val
		l2 = l2.Next
		newNode.Next = mergeTwoLists(l1, l2)
	} else {
		newNode.Val = l1.Val
		l1 = l1.Next
		newNode.Next = mergeTwoLists(l1, l2)
	}
	return newNode
}

func main() {

	// myList := &ListNode{}

	// myList = mergeTwoLists()
	// var i, j int

	// i = 5
	// j = 6
	// 	l1 := ListNode{
	// 		Val: i,

	// 	}

	// 	l2 := ListNode{
	// 		Val: j,
	// 	}

	//mergeTwoLists2(l1,l2)

	//fmt.Println(mergeTwoLists1([]int{1, 1, 2},[]int{1, 1, 2}))

	//l1 = [1,2,4], l2 = [1,3,4]

	// l1 := []int{1,2,4}

	// l2 := []int{1,3,4}

	// sortedList := mergeTwoLists(l1, l2)

	//sortedList := mergeTwoLists([1,2,4], [1,3,4])
}
