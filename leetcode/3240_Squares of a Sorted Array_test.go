package main_test

import (
	"sort"
	"testing"
)

/*
Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.



Example 1:

Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Explanation: After squaring, the array becomes [16,1,0,9,100].
After sorting, it becomes [0,1,9,16,100].
Example 2:

Input: nums = [-7,-3,2,3,11]
Output: [4,9,9,49,121]


Constraints:

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums is sorted in non-decreasing order.


Follow up: Squaring each element and sorting the new array is very trivial, could you find an O(n) solution using a different approach?
*/

func sortedSquares(nums []int) []int {
	result := []int{}
	for _, v := range nums {
		result = append(result, v*v)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return result

}

func Test_3240_Example1(t *testing.T) {
	output := sortedSquares([]int{-4, -1, 0, 3, 10})
	if !intArraysIsEqual(output, []int{0, 1, 9, 16, 100}) {
		t.Fatal()
	}
}

func Test_3240_Example2(t *testing.T) {
	output := sortedSquares([]int{-7, -3, 2, 3, 11})
	if !intArraysIsEqual(output, []int{4, 9, 9, 49, 121}) {
		t.Fatal()
	}
}
