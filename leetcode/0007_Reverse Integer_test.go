package main_test

import (
	"math"
	"testing"
)

/*
7. Reverse Integer
Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).



Example 1:

Input: x = 123
Output: 321
Example 2:

Input: x = -123
Output: -321
Example 3:

Input: x = 120
Output: 21


Constraints:

-231 <= x <= 231 - 1
*/

func reverseInteger(x int) int {

	n := 0
	s := 1

	if x < 0 {
		s = -1
		x = x * -1
	}

	for x > 0 {
		n = n*10 + x%10
		x = x / 10
	}

	if n > math.MaxInt32 || n < (-1)*math.MaxInt32 {
		return 0
	}

	return n * s
}

func Test_0007_Example1(t *testing.T) {
	if reverseInteger(123) != 321 {
		t.Fatal()
	}
}
func Test_0007_Example2(t *testing.T) {
	if reverseInteger(-123) != -321 {
		t.Fatal()
	}
}
func Test_0007_Example3(t *testing.T) {
	if reverseInteger(120) != 21 {
		t.Fatal()
	}
}

func Test_0007_Example5(t *testing.T) {
	if reverseInteger(0) != 0 {
		t.Fatal()
	}
}

func Test_0007_Example6(t *testing.T) {
	if reverseInteger(1534236469) != 0 {
		t.Fatal()
	}
}
