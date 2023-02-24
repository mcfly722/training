package main_test

import (
	"fmt"
	"testing"
)

/*
66. Plus One
You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.

Increment the large integer by one and return the resulting array of digits.



Example 1:

Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Incrementing by one gives 123 + 1 = 124.
Thus, the result should be [1,2,4].
Example 2:

Input: digits = [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
Incrementing by one gives 4321 + 1 = 4322.
Thus, the result should be [4,3,2,2].
Example 3:

Input: digits = [9]
Output: [1,0]
Explanation: The array represents the integer 9.
Incrementing by one gives 9 + 1 = 10.
Thus, the result should be [1,0].


Constraints:

1 <= digits.length <= 100
0 <= digits[i] <= 9
digits does not contain any leading 0's.
*/

func plusOne(digits []int) []int {

	if len(digits) == 0 {
		return []int{1}
	}

	var result []int

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			break
		}
		digits[i] = 0
	}

	if digits[0] == 0 {
		result = append([]int{1}, digits...)
	} else {
		result = digits
	}

	return result
}

func Test_0066_Example1(t *testing.T) {
	result := plusOne([]int{1, 2, 3})
	if !IsEqualIntArrays(result, []int{1, 2, 4}) {
		t.Fatal(fmt.Sprintf("%v", result))
	}
}

func Test_0066_Example2(t *testing.T) {
	result := plusOne([]int{4, 3, 2, 1})
	if !IsEqualIntArrays(result, []int{4, 3, 2, 2}) {
		t.Fatal(fmt.Sprintf("%v", result))
	}
}

func Test_0066_Example4(t *testing.T) {
	result := plusOne([]int{9})
	if !IsEqualIntArrays(result, []int{1, 0}) {
		t.Fatal(fmt.Sprintf("%v", result))
	}
}

func Test_0066_Example5(t *testing.T) {
	result := plusOne([]int{9, 9, 9, 9})
	if !IsEqualIntArrays(result, []int{1, 0, 0, 0, 0}) {
		t.Fatal(fmt.Sprintf("%v", result))
	}
}

func Test_0066_Example6(t *testing.T) {
	result := plusOne([]int{})
	if !IsEqualIntArrays(result, []int{1}) {
		t.Fatal(fmt.Sprintf("%v", result))
	}
}
