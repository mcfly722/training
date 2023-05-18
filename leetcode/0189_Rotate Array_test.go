package leetcode_test

import (
	"fmt"
	"testing"
)

/*
189. Rotate Array
Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

Example 1:

Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]
Example 2:

Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]


Constraints:

1 <= nums.length <= 105
-231 <= nums[i] <= 231 - 1
0 <= k <= 105


Follow up:

Try to come up with as many solutions as you can. There are at least three different ways to solve this problem.
Could you do it in-place with O(1) extra space?
*/

func reverse(nums []int, from int, to int) {
	for i := 0; i <= (to-from)/2; i++ {
		tmp := nums[from+i]
		nums[from+i] = nums[to-i]
		nums[to-i] = tmp
	}
}

// rotate to right
func rotate(nums []int, k int) {
	if len(nums) > 0 {
		k = k % len(nums)

		if k > 0 {
			reverse(nums, 0, len(nums)-1)

			// for right rotation
			reverse(nums, 0, k-1)
			reverse(nums, k, len(nums)-1)

			// for left rotations
			//reverse(nums, 0, len(nums)-1-k)
			//reverse(nums, len(nums)-k, len(nums)-1)

		}
	}
}

func Test_0189_Example0(t *testing.T) {
	sample := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(sample, len(sample))
	if !IsEqualIntArrays(sample, []int{1, 2, 3, 4, 5, 6, 7}) {
		t.Fatalf(fmt.Sprintf("%v", sample))
	}
}

func Test_0189_Example1(t *testing.T) {
	sample := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(sample, 1)
	if !IsEqualIntArrays(sample, []int{7, 1, 2, 3, 4, 5, 6}) {
		t.Fatalf(fmt.Sprintf("%v", sample))
	}
}

func Test_0189_Example3(t *testing.T) {
	sample := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(sample, 3)
	if !IsEqualIntArrays(sample, []int{5, 6, 7, 1, 2, 3, 4}) {
		t.Fatalf(fmt.Sprintf("%v", sample))
	}
}

func Test_0189_Example4(t *testing.T) {
	sample := []int{-1, -100, 3, 99}
	rotate(sample, 2)
	if !IsEqualIntArrays(sample, []int{3, 99, -1, -100}) {
		t.Fatalf(fmt.Sprintf("%v", sample))
	}
}

func Test_0189_Example5(t *testing.T) {
	sample := []int{}
	rotate(sample, 1)
	if !IsEqualIntArrays(sample, []int{}) {
		t.Fatalf(fmt.Sprintf("%v", sample))
	}
}

func Test_0189_Example6(t *testing.T) {
	sample := []int{6}
	rotate(sample, 1)
	if !IsEqualIntArrays(sample, []int{6}) {
		t.Fatalf(fmt.Sprintf("%v", sample))
	}
}
