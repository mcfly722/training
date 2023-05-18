package leetcode_test

import (
	"fmt"
	"testing"
)

/*
75. Sort Colors
Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

You must solve this problem without using the library's sort function.



Example 1:

Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
Example 2:

Input: nums = [2,0,1]
Output: [0,1,2]


Constraints:

n == nums.length
1 <= n <= 300
nums[i] is either 0, 1, or 2.


Follow up: Could you come up with a one-pass algorithm using only constant extra space?

*/

func sortColors(nums []int) {
	red := 0
	white := 0
	blue := 0

	for _, num := range nums {
		switch num {
		case 0:
			red++
		case 1:
			white++
		case 2:
			blue++
		}
	}

	pos := 0

	for i := 0; i < red; i++ {
		nums[pos] = 0
		pos++
	}
	for i := 0; i < white; i++ {
		nums[pos] = 1
		pos++
	}
	for i := 0; i < blue; i++ {
		nums[pos] = 2
		pos++
	}

}

func Test_0075_Example1(t *testing.T) {
	sample := []int{2, 0, 2, 1, 1, 0}
	sortColors(sample)
	if fmt.Sprintf("%v", sample) != "[0 0 1 1 2 2]" {
		t.Fatal(fmt.Sprintf("%v", sample))
	}
}

func Test_0075_Example2(t *testing.T) {
	sample := []int{2, 0, 1}
	sortColors(sample)
	if fmt.Sprintf("%v", sample) != "[0 1 2]" {
		t.Fatal(fmt.Sprintf("%v", sample))
	}
}

func Test_0075_Example3(t *testing.T) {
	sample := []int{}
	sortColors(sample)
	if fmt.Sprintf("%v", sample) != "[]" {
		t.Fatal(fmt.Sprintf("%v", sample))
	}
}
