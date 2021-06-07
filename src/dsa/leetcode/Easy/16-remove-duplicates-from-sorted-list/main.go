package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
We can prove that this condition is indeed a loop invariant by induction.

A loop invariant is condition that is true before and after every iteration of the loop.
 In this case, a loop invariant that helps us prove correctness is this: 

Before going into the loop, current points to the head of the list. 
Therefore, the part of the list up to current contains only the head. 
And so it can not contain any duplicate elements. 
Now suppose current is now pointing to some node in the list (but not the last element), 
and the part of the list up to current contains no duplicate elements. 
After another loop iteration, one of two things happen.
current.next was a duplicate of current. In this case, the duplicate node at current.next is deleted, 
and current stays pointing to the same node as before. 
Therefore, the condition still holds; there are still no duplicates up to current.
current.next was not a duplicate of current (and, because the list is sorted, current.next is also not a duplicate of any other element appearing before current). 
In this case, current moves forward one step to point to current.next. 
Therefore, the condition still holds; there are no duplicates up to current.
At the last iteration of the loop, current must point to the last element, 
because afterwards, current.next = null. Therefore, after the loop ends, 
all elements up to the last element do not contain duplicates.
*/

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	n1 := head
	n2 := head.Next

	for n2 != nil {
//if there will be duplicate value it will get removed as and node will point to next node
		if n2.Val == n1.Val {
			n2 = n2.Next
		} else {
			n1.Next = n2
			n1 = n1.Next
			n2 = n2.Next
		}
	}
	n1.Next = nil
	return head
}
