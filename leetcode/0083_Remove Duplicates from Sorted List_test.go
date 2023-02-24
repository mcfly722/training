package main_test

import (
	"fmt"
	"testing"
)

/*
83. Remove Duplicates from Sorted List
Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

Example 1:

Input: head = [1,1,2]
Output: [1,2]
Example 2:

Input: head = [1,1,2,3,3]
Output: [1,2,3]

Constraints:

The number of nodes in the list is in the range [0, 300].
-100 <= Node.val <= 100
The list is guaranteed to be sorted in ascending order.
*/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	newHead := &ListNode{
		Val:  head.Val,
		Next: nil,
	}

	currentNew := newHead

	for current := head.Next; current != nil; current = current.Next {
		if current.Val > currentNew.Val {
			newNode := &ListNode{
				Val:  current.Val,
				Next: nil,
			}
			currentNew.Next = newNode
			currentNew = newNode
		}
	}

	return newHead
}

func Test_0083_Example1(t *testing.T) {
	result := List2IntArray(deleteDuplicates(IntArray2List([]int{1, 1, 2})))
	if !IsEqual2IntArray(*result, []int{1, 2}) {
		t.Fatal(fmt.Sprintf("%v", result))
	}
}

func Test_0083_Example2(t *testing.T) {
	result := List2IntArray(deleteDuplicates(IntArray2List([]int{1, 1, 2, 3, 3})))
	if !IsEqual2IntArray(*result, []int{1, 2, 3}) {
		t.Fatal(fmt.Sprintf("%v", result))
	}
}

func Test_0083_Example3(t *testing.T) {
	result := List2IntArray(deleteDuplicates(IntArray2List([]int{1, 1})))
	if !IsEqual2IntArray(*result, []int{1}) {
		t.Fatal(fmt.Sprintf("%v", result))
	}
}

func Test_0083_Example4(t *testing.T) {
	result := List2IntArray(deleteDuplicates(IntArray2List([]int{})))
	if !IsEqual2IntArray(*result, []int{}) {
		t.Fatal(fmt.Sprintf("%v", result))
	}
}
