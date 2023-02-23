package main_test

import (
	"math"
	"sort"
	"testing"
)

/*
Given an integer array nums, return the third distinct maximum number in this array. If the third maximum does not exist, return the maximum number.



Example 1:

Input: nums = [3,2,1]
Output: 1
Explanation:
The first distinct maximum is 3.
The second distinct maximum is 2.
The third distinct maximum is 1.
Example 2:

Input: nums = [1,2]
Output: 2
Explanation:
The first distinct maximum is 2.
The second distinct maximum is 1.
The third distinct maximum does not exist, so the maximum (2) is returned instead.
Example 3:

Input: nums = [2,2,3,1]
Output: 1
Explanation:
The first distinct maximum is 3.
The second distinct maximum is 2 (both 2's are counted together since they have the same value).
The third distinct maximum is 1.


Constraints:

1 <= nums.length <= 104
-231 <= nums[i] <= 231 - 1


Follow up: Can you find an O(n) solution?
*/

func sliceContains_3231(arr []int, num int) bool {
	for _, a := range arr {
		if a == num {
			return true
		}
	}
	return false
	//	return slices.Contains(arr, num) // from "golang.org/x/exp/slices"
}

func thirdMax(nums []int) int {
	firstN := []int{}
	for _, current := range nums {
		if !sliceContains_3231(firstN, current) {
			firstN = append(firstN, current)
			sort.Slice(firstN, func(i, j int) bool {
				return firstN[i] >= firstN[j]
			})
			firstN = firstN[:int(math.Min(float64(len(firstN)), 3))]
		}
	}

	//	return firstN[len(firstN)-1]
	if len(firstN) < 3 {
		return firstN[0]
	}
	return firstN[2] // return third
}

func Test_3231_Example1(t *testing.T) {
	thridMax := thirdMax([]int{3, 2, 1})
	if thridMax != 1 {
		t.Fatal(thridMax)
	}
}

func Test_3231_Example2(t *testing.T) {
	thridMax := thirdMax([]int{1, 2})
	if thridMax != 2 {
		t.Fatal(thridMax)
	}
}

func Test_3231_Example3(t *testing.T) {
	thridMax := thirdMax([]int{2, 2, 3, 1})
	if thridMax != 1 {
		t.Fatal(thridMax)
	}
}

func Test_3231_Example4(t *testing.T) {
	thridMax := thirdMax([]int{5, 4, 3, 2, 1})
	if thridMax != 3 {
		t.Fatal(thridMax)
	}
}

func Test_3231_Example5(t *testing.T) {
	thridMax := thirdMax([]int{1, 2})
	if thridMax != 2 {
		t.Fatal(thridMax)
	}
}
