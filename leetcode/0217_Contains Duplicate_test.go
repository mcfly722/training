package main_test

import "testing"

/*
217. Contains Duplicate
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

Example 1:

Input: nums = [1,2,3,1]
Output: true
Example 2:

Input: nums = [1,2,3,4]
Output: false
Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true


Constraints:

1 <= nums.length <= 105
-109 <= nums[i] <= 109
*/

func containsDuplicate(nums []int) bool {
	unique := map[int]bool{}

	for _, num := range nums {
		if _, contain := unique[num]; contain {
			return true
		}
		unique[num] = true
	}

	return false
}

func Test_0217_Example1(t *testing.T) {
	if !containsDuplicate([]int{1, 2, 3, 1}) {
		t.Fatal()
	}
}

func Test_0217_Example2(t *testing.T) {
	if containsDuplicate([]int{1, 2, 3, 4}) {
		t.Fatal()
	}
}

func Test_0217_Example3(t *testing.T) {
	if !containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}) {
		t.Fatal()
	}
}

func Test_0217_Example5(t *testing.T) {
	if containsDuplicate([]int{1}) {
		t.Fatal()
	}
}

func Test_0217_Example6(t *testing.T) {
	if containsDuplicate([]int{}) {
		t.Fatal()
	}
}
