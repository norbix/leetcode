package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{} // Dummy node to simplify handling the merged list.
	current := dummy     // Pointer to the current node in the merged list.

	// Iterate through both lists while either list is not empty.
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1 // Add the node from list1 to the merged list.
			list1 = list1.Next
		} else {
			current.Next = list2 // Add the node from list2 to the merged list.
			list2 = list2.Next
		}
		current = current.Next // Move the current pointer to the last added node.
	}

	// If either list is not empty, add the remaining nodes to the merged list.
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return dummy.Next
}

func main() {
	input1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	input2 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	result := MergeTwoLists(&input1, &input2)

	for result != nil {
		fmt.Printf("%d, ", result.Val)
		result = result.Next
	}
}
