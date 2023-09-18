package main

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	input1, input2 := generatorInputFixtures()
	expected := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
		},
	}
	result := MergeTwoLists(&input1, &input2)

	reflect.DeepEqual(expected, result)
}

func generatorInputFixtures() (ListNode, ListNode) {
	return ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
		ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		}
}
