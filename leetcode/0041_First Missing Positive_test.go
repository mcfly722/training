package main_test

import (
	"math"
	"testing"
)

/*

41. First Missing Positive
Given an unsorted integer array nums, return the smallest missing positive integer.

You must implement an algorithm that runs in O(n) time and uses constant extra space.



Example 1:

Input: nums = [1,2,0]
Output: 3
Explanation: The numbers in the range [1,2] are all in the array.
Example 2:

Input: nums = [3,4,-1,1]
Output: 2
Explanation: 1 is in the array but 2 is missing.
Example 3:

Input: nums = [7,8,9,11,12]
Output: 1
Explanation: The smallest positive integer 1 is missing.


Constraints:

1 <= nums.length <= 105
-231 <= nums[i] <= 231 - 1

*/

func firstMissingPositive(nums []int) int {

	//fmt.Println(fmt.Sprintf("orig: %v", nums))

	// mark all negative numbers and 0 as maxInt32 number

	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			nums[i] = math.MaxInt32
		}
	}

	//fmt.Println(fmt.Sprintf("zeros: %v", nums))

	// mark all existing indexes with negative flag,
	for i := 0; i < len(nums); i++ {

		pos := int(math.Abs(float64(nums[i]))) - 1

		if pos >= 0 && pos < len(nums) {
			if nums[pos] > 0 {
				nums[pos] = nums[pos] * -1
			}
		}
	}

	//fmt.Println(fmt.Sprintf("neg: %v", nums))

	// find first positive - it is missed number
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}

	return len(nums) + 1
}

func Test_0041_Example1(t *testing.T) {
	result := firstMissingPositive([]int{1, 2, 0})
	if result != 3 {
		t.Fatal(result)
	}
}

func Test_0041_Example2(t *testing.T) {
	result := firstMissingPositive([]int{3, 4, -1, 1})
	if result != 2 {
		t.Fatal(result)
	}
}

func Test_0041_Example3(t *testing.T) {
	result := firstMissingPositive([]int{7, 8, 9, 11, 12})
	if result != 1 {
		t.Fatal(result)
	}
}

func Test_0041_Example4(t *testing.T) {
	result := firstMissingPositive([]int{0})
	if result != 1 {
		t.Fatal(result)
	}
}

func Test_0041_Example5(t *testing.T) {
	result := firstMissingPositive([]int{-1})
	if result != 1 {
		t.Fatal(result)
	}
}

func Test_0041_Example6(t *testing.T) {
	result := firstMissingPositive([]int{0, 2})
	if result != 1 {
		t.Fatal(result)
	}
}

func Test_0041_Example7(t *testing.T) {
	result := firstMissingPositive([]int{1})
	if result != 2 {
		t.Fatal(result)
	}
}
