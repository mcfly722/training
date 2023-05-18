package leetcode_test

import (
	"fmt"
	"sort"
	"testing"
)

/*
15. 3Sum
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.



Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.
Example 2:

Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.
Example 3:

Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.


Constraints:

3 <= nums.length <= 3000
-105 <= nums[i] <= 105

*/

/*
--------------------------------------------------------------------------------------------------------------------------------------------------------
// Not optimal solution

func threeSum(nums []int) [][]int {
	result := [][]int{}

	for i := 0; i+2 < len(nums); i++ {
		for j := i + 1; j+1 < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					result = append(result, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}

	return result
}

*/

func threeSum(nums []int) [][]int {
	distinct := map[int64]bool{}

	result := [][]int{}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i+2 < len(nums); i++ {

		left := i + 1
		right := len(nums) - 1

		for left < right {

			if nums[i]+nums[left]+nums[right] == 0 {
				cantor := CantorPairingFunctionWithNegativeAndCommutativeForThree(nums[i], nums[left], nums[right])

				if _, found := distinct[cantor]; !found {
					result = append(result, []int{nums[i], nums[left], nums[right]})
					distinct[cantor] = true
				}

				for ; left < right && nums[left] == nums[left+1]; left++ {
				}

				for ; right > i && nums[right] == nums[right-1]; right-- {
				}

				left++
				right--
			} else {
				if nums[i]+nums[left]+nums[right] < 0 {
					left++
				} else {
					right--
				}

			}

		}

	}

	return result
}

func Test_0015_Example0(t *testing.T) {
	result := threeSum([]int{})
	if len(result) != 0 {
		t.Fatal()
	}
}

func Test_0015_Example1(t *testing.T) {
	result := threeSum([]int{0})
	if len(result) != 0 {
		t.Fatal()
	}
}

func Test_0015_Example2(t *testing.T) {
	result := fmt.Sprintf("%v", threeSum([]int{0, 0, 0}))
	if result != "[[0 0 0]]" {
		t.Fatal(result)
	}
}

func Test_0015_Example3(t *testing.T) {
	result := fmt.Sprintf("%v", threeSum([]int{-1, 0, 1, 2, -1, -4}))
	if result != "[[-1 -1 2] [-1 0 1]]" {
		t.Fatal(result)
	}
}

func Test_0015_Example4(t *testing.T) {
	result := fmt.Sprintf("%v", threeSum([]int{0, 0, 0, 0}))
	if result != "[[0 0 0]]" {
		t.Fatal(result)
	}
}

func Test_0015_Example5(t *testing.T) {
	result := fmt.Sprintf("%v", threeSum([]int{-2, 0, 1, 1, 2}))
	if result != "[[-2 0 2] [-2 1 1]]" {
		t.Fatal(result)
	}
}

func Test_0015_Example6(t *testing.T) {
	result := fmt.Sprintf("%v", threeSum([]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}))
	if result != "[[-4 0 4] [-4 1 3] [-3 -1 4] [-3 0 3] [-3 1 2] [-2 -1 3] [-2 0 2] [-1 -1 2] [-1 0 1]]" {
		t.Fatal(result)
	}
}
