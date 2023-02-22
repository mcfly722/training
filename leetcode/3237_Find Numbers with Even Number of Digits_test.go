package main_test

import (
	"testing"
)

/*
Given an array nums of integers, return how many of them contain an even number of digits.



Example 1:

Input: nums = [12,345,2,6,7896]
Output: 2
Explanation:
12 contains 2 digits (even number of digits).
345 contains 3 digits (odd number of digits).
2 contains 1 digit (odd number of digits).
6 contains 1 digit (odd number of digits).
7896 contains 4 digits (even number of digits).
Therefore only 12 and 7896 contain an even number of digits.
Example 2:

Input: nums = [555,901,482,1771]
Output: 1
Explanation:
Only 1771 contains an even number of digits.


Constraints:

1 <= nums.length <= 500
1 <= nums[i] <= 105
   Hide Hint #1
How to compute the number of digits of a number ?
   Hide Hint #2
Divide the number by 10 again and again to get the number of digits.
*/

func NumberOfDigits(number int) int {
	if number == 0 {
		return 1
	}

	c := 0

	for ; number > 0; c++ {
		number = number / 10
	}

	return c
}

func findNumbers(nums []int) int {
	counter := 0
	for _, v := range nums {
		if NumberOfDigits(v)%2 == 0 {
			counter++
		}
	}

	return counter
}

func Test_3237_Example1(t *testing.T) {
	if findNumbers([]int{12, 345, 2, 6, 7896}) != 2 {
		t.Fatal()
	}
}

func Test_3237_Example_Number_0(t *testing.T) {
	if NumberOfDigits(0) != 1 {
		t.Fatal()
	}
}

func Test_3237_Example_Number_1(t *testing.T) {
	if NumberOfDigits(1) != 1 {
		t.Fatal()
	}
}

func Test_3237_Example_Number_123(t *testing.T) {
	if NumberOfDigits(123) != 3 {
		t.Fatal()
	}
}

func Test_3237_Example_Number_1234(t *testing.T) {
	if NumberOfDigits(1234) != 4 {
		t.Fatal()
	}
}
