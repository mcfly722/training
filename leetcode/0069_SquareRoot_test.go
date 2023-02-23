package main_test

import (
	"testing"
)

/*
69. Sqrt(x)

Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.

You must not use any built-in exponent function or operator.

For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.


Example 1:

Input: x = 4
Output: 2
Explanation: The square root of 4 is 2, so we return 2.
Example 2:

Input: x = 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since we round it down to the nearest integer, 2 is returned.


Constraints:

0 <= x <= 231 - 1

*/

func mySqrt(x int) int {

	if x == 1 {
		return 1
	}

	lower := 0
	upper := x

	for {
		p := (lower + upper) / 2

		if p*p > x {
			upper = p
		}

		if p*p < x {
			lower = p
		}

		if p*p == x {
			return p
		}

		if lower+1 == upper {
			return lower
		}

	}
}

func Test_0069_Example1(t *testing.T) {
	root := mySqrt(4)
	if root != 2 {
		t.Fatal(root)
	}
}

func Test_0069_Example2(t *testing.T) {
	root := mySqrt(8)
	if root != 2 {
		t.Fatal(root)
	}
}

func Test_0069_Example3(t *testing.T) {
	root := mySqrt(1)
	if root != 1 {
		t.Fatal(root)
	}
}

func Test_0069_Example4(t *testing.T) {
	root := mySqrt(0)
	if root != 0 {
		t.Fatal(root)
	}
}
