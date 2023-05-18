package leetcode_test

import (
	"fmt"
	"testing"
)

/*
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.



Example 1:

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
Example 2:

Input: nums = [0]
Output: [0]


Constraints:

1 <= nums.length <= 104
-231 <= nums[i] <= 231 - 1


Follow up: Could you minimize the total number of operations done?
   Hide Hint #1
In-place means we should not be allocating any space for extra array. But we are allowed to modify the existing array. However, as a first step, try coming up with a solution that makes use of additional space. For this problem as well, first apply the idea discussed using an additional array and the in-place solution will pop up eventually.
   Hide Hint #2
A two-pointer approach could be helpful here. The idea would be to have one pointer for iterating the array and another pointer that just works on the non-zero elements of the array.

*/

func shiftIntArrayLeft_3157(nums []int, pos int) {
	for i := pos + 1; i <= len(nums)-1; i++ {
		nums[i-1] = nums[i]
	}
}

func moveZeroes(nums []int) {
	shifts := 0
	for i := 0; i <= len(nums)-1; i++ {
		if nums[i] == 0 && i < len(nums)-shifts {
			shiftIntArrayLeft_3157(nums, i)
			nums[len(nums)-1] = 0
			shifts++
			i--
		}
	}
}

func Test_3157_Example1(t *testing.T) {
	nums := []int{0, 1, 2, 3}

	shiftIntArrayLeft_3247(nums, 2)
	nums = nums[:len(nums)-1]

	if !IsEqualIntArrays(nums, []int{0, 1, 3}) {
		t.Fatalf(fmt.Sprintf("nums=%v", nums))
	}
}

func Test_3157_Example2(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)

	if !IsEqualIntArrays(nums, []int{1, 3, 12, 0, 0}) {
		t.Fatalf(fmt.Sprintf("%v", nums))
	}
}

func Test_3157_Example3(t *testing.T) {
	nums := []int{0}
	moveZeroes(nums)

	if !IsEqualIntArrays(nums, []int{0}) {
		t.Fatalf(fmt.Sprintf("%v", nums))
	}
}

func Test_3157_Example4(t *testing.T) {
	nums := []int{0, 0, 1}
	moveZeroes(nums)

	if !IsEqualIntArrays(nums, []int{1, 0, 0}) {
		t.Fatalf(fmt.Sprintf("%v", nums))
	}
}
