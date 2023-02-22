package main_test

import "testing"

/*
Given an array of integers arr, return true if and only if it is a valid mountain array.

Recall that arr is a mountain array if and only if:

arr.length >= 3
There exists some i with 0 < i < arr.length - 1 such that:
arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
arr[i] > arr[i + 1] > ... > arr[arr.length - 1]



Example 1:

Input: arr = [2,1]
Output: false
Example 2:

Input: arr = [3,5,5]
Output: false
Example 3:

Input: arr = [0,3,2,1]
Output: true


Constraints:

1 <= arr.length <= 104
0 <= arr[i] <= 104
   Hide Hint #1
It's very easy to keep track of a monotonically increasing or decreasing ordering of elements. You just need to be able to determine the start of the valley in the mountain and from that point onwards, it should be a valley i.e. no mini-hills after that. Use this information in regards to the values in the array and you will be able to come up with a straightforward solution.
*/

func validMountainArray(arr []int) bool {
	if len(arr) >= 3 {
		inc_index := 0
		dec_index := len(arr) - 1

		inc_value := arr[inc_index] - 1
		dec_value := arr[dec_index] - 1

		for ; inc_index < len(arr)-1 && arr[inc_index] > inc_value; inc_index++ {
			inc_value = arr[inc_index]
		}

		for ; dec_index > 0 && arr[dec_index] > dec_value; dec_index-- {
			dec_value = arr[dec_index]
		}

		if inc_index-1 == dec_index+1 {
			return true
		}
	}

	return false
}

func Test_3251_Example1(t *testing.T) {
	if validMountainArray([]int{2, 1}) {
		t.Fatal()
	}
}

func Test_3251_Example2(t *testing.T) {
	if validMountainArray([]int{3, 5, 5}) {
		t.Fatal()
	}
}

func Test_3251_Example3(t *testing.T) {
	if !validMountainArray([]int{0, 3, 2, 1}) {
		t.Fatal()
	}
}

func Test_3251_Example4(t *testing.T) {
	if !validMountainArray([]int{0, 2, 3, 4, 5, 2, 1, 0}) {
		t.Fatal()
	}
}

func Test_3251_Example5(t *testing.T) {
	if validMountainArray([]int{0, 2, 3, 3, 5, 2, 1, 1}) {
		t.Fatal()
	}
}
