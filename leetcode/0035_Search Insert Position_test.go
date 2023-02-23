package main_test

import "testing"

/*
35. Search Insert Position
Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.



Example 1:

Input: nums = [1,3,5,6], target = 5
Output: 2
Example 2:

Input: nums = [1,3,5,6], target = 2
Output: 1
Example 3:

Input: nums = [1,3,5,6], target = 7
Output: 4


Constraints:

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums contains distinct values sorted in ascending order.
-104 <= target <= 104
*/

func searchInsert(nums []int, target int) int {

	if len(nums) == 0 {
		return 0
	}

	if target <= nums[0] {
		return 0
	}

	if target > nums[len(nums)-1] {
		return len(nums)
	}

	left := 0
	right := len(nums) - 1

	for {
		p := (left + right) / 2

		if target > nums[p] {
			left = p
		}

		if target <= nums[p] {
			right = p
		}

		if left == right-1 {

			return right
		}

	}
}

func Test_0035_Example1(t *testing.T) {
	result := searchInsert([]int{1, 3, 5, 6}, 5)
	if !(result == 2) {
		t.Fatal(result)
	}
}

func Test_0035_Example2(t *testing.T) {
	result := searchInsert([]int{1, 3, 5, 6}, 2)
	if result != 1 {
		t.Fatal(result)
	}
}

func Test_0035_Example3(t *testing.T) {
	result := searchInsert([]int{1, 3, 5, 6}, 7)
	if result != 4 {
		t.Fatal(result)
	}
}

func Test_0035_Example4(t *testing.T) {
	result := searchInsert([]int{1, 3, 5, 6}, -1)
	if result != 0 {
		t.Fatal(result)
	}
}

func Test_0035_Example5(t *testing.T) {
	result := searchInsert([]int{}, 5)
	if result != 0 {
		t.Fatal(result)
	}
}

func Test_0035_Example6(t *testing.T) {
	result := searchInsert([]int{1, 3}, 3)
	if result != 1 {
		t.Fatal(result)
	}
}
