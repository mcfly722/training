package main_test

import (
	"fmt"
	"testing"
)

/*
46. Permutations
Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.

Example 1:

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
Example 2:

Input: nums = [0,1]
Output: [[0,1],[1,0]]
Example 3:

Input: nums = [1]
Output: [[1]]

Constraints:

1 <= nums.length <= 6
-10 <= nums[i] <= 10
All the integers of nums are unique.

*/

func copyWithoutElement(slice []int, i int) []int {
	new := []int{}
	new = append(new, slice[:i]...)
	new = append(new, slice[i+1:]...)
	return new
}

func permute(nums []int) [][]int {

	//fmt.Println(fmt.Sprintf("%v", nums))

	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	if len(nums) == 2 {
		return [][]int{{nums[0], nums[1]}, {nums[1], nums[0]}}
	}

	resultedPermutations := [][]int{}

	for excludeIndex := 0; excludeIndex < len(nums); excludeIndex++ {

		excludedNumber := nums[excludeIndex]
		leftNumbers := copyWithoutElement(nums, excludeIndex)
		//		fmt.Println(fmt.Sprintf("original:%v excludedIndex:%v leftNumbers:%v", nums, excludeIndex, leftNumbers))
		newPermutations := permute(leftNumbers)

		for _, newPermutation := range newPermutations {
			resultedPermutation := append([]int{excludedNumber}, newPermutation...)
			//fmt.Println(fmt.Sprintf("resulted:%v", resultedPermutation))
			resultedPermutations = append(resultedPermutations, resultedPermutation)
		}

	}

	return resultedPermutations
}

func Test_0046_Example1(t *testing.T) {
	result := fmt.Sprintf("%v", permute([]int{1, 2, 3}))
	if result != "[[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]" {
		t.Fatal(result)
	}
}

func Test_0046_Example2(t *testing.T) {
	result := fmt.Sprintf("%v", permute([]int{1, 2}))
	if result != "[[1 2] [2 1]]" {
		t.Fatal(result)
	}
}

func Test_0046_Example3(t *testing.T) {
	result := fmt.Sprintf("%v", permute([]int{1}))
	if result != "[[1]]" {
		t.Fatal(result)
	}
}

func Test_0046_Example4(t *testing.T) {
	result := fmt.Sprintf("%v", permute([]int{}))
	if result != "[]" {
		t.Fatal(result)
	}
}
