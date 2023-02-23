package main_test

import (
	"fmt"
	"testing"
)

/*
Given an integer array nums, move all the even integers at the beginning of the array followed by all the odd integers.

Return any array that satisfies this condition.



Example 1:

Input: nums = [3,1,2,4]
Output: [2,4,3,1]
Explanation: The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.
Example 2:

Input: nums = [0]
Output: [0]


Constraints:

1 <= nums.length <= 5000
0 <= nums[i] <= 5000

*/

func shiftIntArrayLeft_3260(nums []int, pos int) {
	for i := pos + 1; i <= len(nums)-1; i++ {
		nums[i-1] = nums[i]
	}
}

func sortArrayByParity(nums []int) []int {
	result := nums
	moves := 0

	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 1 && i < len(nums)-moves {
			tmp := nums[i]
			shiftIntArrayLeft_3260(nums, i)
			nums[len(nums)-1] = tmp
			i--
			moves++
		}
	}
	return result
}

func Test_3260_shiftIntArrayLeft1(t *testing.T) {
	nums := []int{0, 1, 2, 3}

	shiftIntArrayLeft_3260(nums, 2)
	nums = nums[:len(nums)-1]

	if !intArraysIsEqual(nums, []int{0, 1, 3}) {
		t.Fatal(fmt.Sprintf("nums=%v", nums))
	}
}

func Test_3260_Example1(t *testing.T) {
	input := []int{3, 1, 2, 4}

	output := sortArrayByParity(input)

	if !(intArraysIsEqual(output, []int{2, 4, 3, 1}) || intArraysIsEqual(output, []int{2, 4, 1, 3}) || intArraysIsEqual(output, []int{4, 2, 1, 3})) {
		t.Fatal(fmt.Sprintf("%v", output))
	}
}

func Test_3260_Example2(t *testing.T) {
	input := []int{0}

	output := sortArrayByParity(input)

	if !(intArraysIsEqual(output, []int{0})) {
		t.Fatal(fmt.Sprintf("%v", output))
	}
}
