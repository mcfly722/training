package leetcode_test

import (
	"fmt"
	"testing"
)

/*
1. Two Sum
Easy
43.6K
1.4K
Companies
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.



Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]


Constraints:

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.


Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?
*/

func twoSum(nums []int, target int) []int {
	index := map[int]int{}
	resultMap := map[int]int{}

	result := []int{}

	for i, v := range nums {
		index[v] = i
	}

	for i1, v := range nums {
		if i2, ok := index[target-v]; ok {
			if i1 != i2 {
				resultMap[i1] = 0
				resultMap[i2] = 0
			}
		}
	}

	for k := range resultMap {
		result = append(result, k)
	}

	return result
}

func Test_0001_Example1(t *testing.T) {
	sumsIndex := twoSum([]int{2, 7, 11, 15}, 9)
	if !(IsEqualIntArrays(sumsIndex, []int{0, 1}) || IsEqualIntArrays(sumsIndex, []int{1, 0})) {
		t.Fatal(fmt.Sprintf("%v", sumsIndex))
	}
}

func Test_0001_Example2(t *testing.T) {
	sumsIndex := twoSum([]int{3, 2, 4}, 6)
	if !(IsEqualIntArrays(sumsIndex, []int{1, 2}) || IsEqualIntArrays(sumsIndex, []int{2, 1})) {
		t.Fatal(fmt.Sprintf("%v", sumsIndex))
	}
}

func Test_0001_Example3(t *testing.T) {
	sumsIndex := twoSum([]int{3, 3}, 6)
	if !(IsEqualIntArrays(sumsIndex, []int{0, 1}) || IsEqualIntArrays(sumsIndex, []int{1, 0})) {
		t.Fatal(fmt.Sprintf("%v", sumsIndex))
	}
}
