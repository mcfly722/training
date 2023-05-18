package leetcode_test

import (
	"fmt"
	"testing"
)

/*
1480. Running Sum of 1d Array
Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).

Return the running sum of nums.



Example 1:

Input: nums = [1,2,3,4]
Output: [1,3,6,10]
Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].
Example 2:

Input: nums = [1,1,1,1,1]
Output: [1,2,3,4,5]
Explanation: Running sum is obtained as follows: [1, 1+1, 1+1+1, 1+1+1+1, 1+1+1+1+1].
Example 3:

Input: nums = [3,1,2,10,1]
Output: [3,4,6,16,17]


Constraints:

1 <= nums.length <= 1000
-10^6 <= nums[i] <= 10^6

*/

func runningSum(nums []int) []int {
	result := []int{}

	counter := 0

	for _, num := range nums {
		result = append(result, counter+num)
		counter += num
	}

	return result
}

func Test_1480_Example1(t *testing.T) {
	output := runningSum([]int{1, 2, 3, 4})
	if !IsEqualIntArrays(output, []int{1, 3, 6, 10}) {
		t.Fatalf(fmt.Sprintf("%v", output))
	}
}

func Test_1480_Example2(t *testing.T) {
	output := runningSum([]int{1, 1, 1, 1, 1})
	if !IsEqualIntArrays(output, []int{1, 2, 3, 4, 5}) {
		t.Fatalf(fmt.Sprintf("%v", output))
	}
}

func Test_1480_Example4(t *testing.T) {
	output := runningSum([]int{3, 1, 2, 10, 1})
	if !IsEqualIntArrays(output, []int{3, 4, 6, 16, 17}) {
		t.Fatalf(fmt.Sprintf("%v", output))
	}
}

func Test_1480_Example5(t *testing.T) {
	output := runningSum([]int{})
	if !IsEqualIntArrays(output, []int{}) {
		t.Fatalf(fmt.Sprintf("%v", output))
	}
}

func Test_1480_Example6(t *testing.T) {
	output := runningSum([]int{5})
	if !IsEqualIntArrays(output, []int{5}) {
		t.Fatalf(fmt.Sprintf("%v", output))
	}
}
