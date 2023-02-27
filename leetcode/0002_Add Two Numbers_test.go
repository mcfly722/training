package main_test

import (
	"fmt"
	"math"
	"testing"
)

/*
2. Add Two Numbers
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.



Example 1:


Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]
Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]


Constraints:

The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.

*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode = nil
	var tail *ListNode = nil

	remainder := 0

	var first *ListNode = l1
	var second *ListNode = l2

	for first != nil || second != nil || remainder != 0 {

		if first != nil {
			//			fmt.Print(fmt.Sprintf("%v", first.Val))
			remainder += first.Val
		}
		if second != nil {
			//			fmt.Print(fmt.Sprintf(",%v", second.Val))
			remainder += second.Val
		}
		//		fmt.Println(fmt.Sprintf("->%v", remainder))

		node := &ListNode{
			Val:  remainder % 10,
			Next: nil,
		}

		remainder = (int)(math.Floor((float64)(remainder) / 10))

		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}

		if first != nil {
			first = first.Next
		}

		if second != nil {
			second = second.Next
		}

	}

	return head
}

func Test_0002_Example1(t *testing.T) {
	l1 := IntArray2List([]int{2, 4, 3})
	l2 := IntArray2List([]int{5, 6, 4})

	result := List2IntArray(addTwoNumbers(l1, l2))
	if !IsEqualIntArrays(*result, []int{7, 0, 8}) {
		t.Fatal(fmt.Sprintf("%v", result))
	}
}

func Test_0002_Example2(t *testing.T) {
	l1 := IntArray2List([]int{0})
	l2 := IntArray2List([]int{0})

	result := List2IntArray(addTwoNumbers(l1, l2))
	if !IsEqualIntArrays(*result, []int{0}) {
		t.Fatal(fmt.Sprintf("%v", result))
	}
}

func Test_0002_Example3(t *testing.T) {
	l1 := IntArray2List([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := IntArray2List([]int{9, 9, 9, 9})

	result := List2IntArray(addTwoNumbers(l1, l2))
	if !IsEqualIntArrays(*result, []int{8, 9, 9, 9, 0, 0, 0, 1}) {
		t.Fatal(fmt.Sprintf("%v", result))
	}
}

func Test_0002_Example4(t *testing.T) {
	l1 := IntArray2List([]int{})
	l2 := IntArray2List([]int{})

	result := List2IntArray(addTwoNumbers(l1, l2))
	if !IsEqualIntArrays(*result, []int{}) {
		t.Fatal(fmt.Sprintf("%v", result))
	}
}
