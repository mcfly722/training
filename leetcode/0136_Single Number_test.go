package leetcode_test

import "testing"

/*
136. Single Number
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.



Example 1:

Input: nums = [2,2,1]
Output: 1
Example 2:

Input: nums = [4,1,2,1,2]
Output: 4
Example 3:

Input: nums = [1]
Output: 1


Constraints:

1 <= nums.length <= 3 * 104
-3 * 104 <= nums[i] <= 3 * 104
Each element in the array appears twice except for one element which appears only once.
*/

func singleNumber(nums []int) int {

	n := 0

	for _, num := range nums {
		n = num ^ n // xor
	}

	return n
}

func Test_0136_Example1(t *testing.T) {
	result := singleNumber([]int{2, 2, 1})
	if result != 1 {
		t.Fatal()
	}
}

func Test_0136_Example2(t *testing.T) {
	result := singleNumber([]int{4, 1, 2, 1, 2})
	if result != 4 {
		t.Fatal()
	}
}

func Test_0136_Example3(t *testing.T) {
	result := singleNumber([]int{1})
	if result != 1 {
		t.Fatal()
	}
}
