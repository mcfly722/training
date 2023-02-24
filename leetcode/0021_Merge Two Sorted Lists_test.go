package main_test

import (
	"fmt"
	"testing"
)

/*
21. Merge Two Sorted Lists
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.



Example 1:


Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
Example 2:

Input: list1 = [], list2 = []
Output: []
Example 3:

Input: list1 = [], list2 = [0]
Output: [0]


Constraints:

The number of nodes in both lists is in the range [0, 50].
-100 <= Node.val <= 100
Both list1 and list2 are sorted in non-decreasing order.

*/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var head *ListNode = list1

	if list1.Val > list2.Val {
		head = list2
		list2 = list2.Next
	} else {
		list1 = list1.Next
	}

	var current *ListNode = head

	for list1 != nil && list2 != nil {

		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 == nil {
		current.Next = list2
	} else {
		current.Next = list1
	}

	return head
}

func Test_0021_Example1(t *testing.T) {
	l1 := IntArray2List([]int{1, 2, 4})
	l2 := IntArray2List([]int{1, 3, 4})

	result := List2IntArray(mergeTwoLists(l1, l2))

	if !intArraysIsEqual(*result, []int{1, 1, 2, 3, 4, 4}) {
		t.Fatal(fmt.Sprintf("%v", result))
	}
}

func Test_0021_Example2(t *testing.T) {
	l1 := IntArray2List([]int{})
	l2 := IntArray2List([]int{1, 3, 4})

	result := List2IntArray(mergeTwoLists(l1, l2))

	if !intArraysIsEqual(*result, []int{1, 3, 4}) {
		t.Fatal(fmt.Sprintf("%v", result))
	}
}

func Test_0021_Example3(t *testing.T) {
	l1 := IntArray2List([]int{1, 2, 4})
	l2 := IntArray2List([]int{})

	result := List2IntArray(mergeTwoLists(l1, l2))

	if !intArraysIsEqual(*result, []int{1, 2, 4}) {
		t.Fatal(fmt.Sprintf("%v", result))
	}
}

func Test_0021_Example4(t *testing.T) {
	l1 := IntArray2List([]int{})
	l2 := IntArray2List([]int{})

	result := List2IntArray(mergeTwoLists(l1, l2))

	if !intArraysIsEqual(*result, []int{}) {
		t.Fatal(fmt.Sprintf("%v", result))
	}
}
