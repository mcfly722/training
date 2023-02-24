package main_test

import (
	"fmt"
	"math"
	"testing"
)

/*
Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the range [1, n] that do not appear in nums.



Example 1:

Input: nums = [4,3,2,7,8,2,3,1]
Output: [5,6]
Example 2:

Input: nums = [1,1]
Output: [2]


Constraints:

n == nums.length
1 <= n <= 105
1 <= nums[i] <= n


Follow up: Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

   Hide Hint #1
This is a really easy problem if you decide to use additional memory. For those trying to write an initial solution using additional memory, think counters!
   Hide Hint #2
However, the trick really is to not use any additional space than what is already available to use. Sometimes, multiple passes over the input array help find the solution. However, there's an interesting piece of information in this problem that makes it easy to re-use the input array itself for the solution.
   Hide Hint #3
The problem specifies that the numbers in the array will be in the range [1, n] where n is the number of elements in the array. Can we use this information and modify the array in-place somehow to find what we need?

*/

func sliceContains_3270(arr []int, num int) bool {
	for _, a := range arr {
		if a == num {
			return true
		}
	}
	return false
	//	return slices.Contains(arr, num) // from "golang.org/x/exp/slices"
}

func findDisappearedNumbers(nums []int) []int {

	//mark all appears as negative
	for i := 0; i < len(nums); i++ {

		v := int(math.Abs(float64(nums[i])) - 1)
		if nums[v] > 0 {
			nums[v] = -1 * nums[v]
		}
	}

	// get all indicies with positive numbers
	result := []int{}

	for i, v := range nums {
		if v > 0 {
			result = append(result, i+1)
		}
	}

	return result
}

func Test_3270_Example1(t *testing.T) {
	disappearedNumbers := findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})
	if !IsEqual2IntArray(disappearedNumbers, []int{5, 6}) {
		t.Fatal(fmt.Sprintf("%v", disappearedNumbers))
	}
}

func Test_3270_Example2(t *testing.T) {
	disappearedNumbers := findDisappearedNumbers([]int{1, 1})
	if !IsEqual2IntArray(disappearedNumbers, []int{2}) {
		t.Fatal(fmt.Sprintf("%v", disappearedNumbers))
	}
}
