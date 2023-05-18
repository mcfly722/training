package leetcode_test

import (
	"fmt"
	"testing"
)

func middleNode(head *ListNode) *ListNode {
	last := head
	first := head

	i := 0
	for ; first.Next != nil; first = first.Next {
		if i%2 == 0 {
			last = last.Next
		}
		i++
	}

	return last
}

func Test_0876_Example1(t *testing.T) {
	result := List2IntArray(middleNode(IntArray2List([]int{1, 2, 3, 4, 5})))
	if !(IsEqualIntArrays(*result, []int{3, 4, 5})) {
		t.Fatalf(fmt.Sprintf("%v", result))
	}

}

func Test_0876_Example2(t *testing.T) {
	result := List2IntArray(middleNode(IntArray2List([]int{1, 2, 3, 4, 5, 6})))
	if !(IsEqualIntArrays(*result, []int{4, 5, 6})) {
		t.Fatalf(fmt.Sprintf("%v", result))
	}
}
