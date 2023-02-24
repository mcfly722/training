package main_test

type ListNode struct {
	Val  int
	Next *ListNode
}

func IntArray2List(arr []int) *ListNode {

	if len(arr) == 0 {
		return nil
	}

	var first *ListNode = nil
	var current *ListNode = nil

	for _, element := range arr {

		node := &ListNode{
			Val:  element,
			Next: nil,
		}

		if current != nil {
			current.Next = node
		}

		if first == nil {
			first = node
		}

		current = node

	}

	return first
}

func list2Arr(head *ListNode) *[]int {
	result := []int{}

	if head == nil {
		return &result
	}

	var current *ListNode = head

	for {
		result = append(result, current.Val)

		if current.Next == nil {
			return &result
		}
		current = current.Next
	}

}
