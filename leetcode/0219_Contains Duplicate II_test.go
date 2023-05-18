package leetcode_test

import (
	"testing"
)

/*
219. Contains Duplicate II
Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.

Example 1:

Input: nums = [1,2,3,1], k = 3
Output: true
Example 2:

Input: nums = [1,0,1,1], k = 1
Output: true
Example 3:

Input: nums = [1,2,3,1,2,3], k = 2
Output: false


Constraints:

1 <= nums.length <= 105
-109 <= nums[i] <= 109
0 <= k <= 105
*/

func containsNearbyDuplicate(nums []int, k int) bool {
	dupes := map[int]int{}

	for i, num := range nums {
		if _, contain := dupes[num]; contain {
			if i-dupes[num] <= k {
				return true
			}
		}
		dupes[num] = i
	}

	return false
}

func Test_0219_Example0(t *testing.T) {

	if containsNearbyDuplicate([]int{}, 10) {
		t.Fatal()
	}
}

func Test_0219_Example1(t *testing.T) {
	if !containsNearbyDuplicate([]int{1, 2, 3, 1}, 3) {
		t.Fatal()
	}
}

func Test_0219_Example2(t *testing.T) {
	if !containsNearbyDuplicate([]int{1, 0, 1, 1}, 1) {
		t.Fatal()
	}
}

func Test_0219_Example3(t *testing.T) {
	if containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2) {
		t.Fatal()
	}
}
