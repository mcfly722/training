package leetcode_test

import (
	"fmt"
	"math"
	"testing"
)

/*
50. Pow(x, n)

Implement pow(x, n), which calculates x raised to the power n (i.e., xn).

Example 1:

Input: x = 2.00000, n = 10
Output: 1024.00000
Example 2:

Input: x = 2.10000, n = 3
Output: 9.26100
Example 3:

Input: x = 2.00000, n = -2
Output: 0.25000
Explanation: 2-2 = 1/22 = 1/4 = 0.25


Constraints:

-100.0 < x < 100.0
-231 <= n <= 231-1
n is an integer.
-104 <= xn <= 104

*/

// how many times it would iterate over multiplication
func pow2(x float64, p int) float64 {
	if p == 0 {
		return x
	}
	for i := 0; i < p; i++ {
		x = x * x
	}
	return x
}

func myPow(x float64, n int) float64 {

	c := int(math.Abs(float64(n)))

	result := float64(1)

	for j := 0; c > 0; j++ {
		v := c % 2
		c = c / 2

		if v != 0 {
			result = result * pow2(x, j)
		}
	}

	if n < 0 {
		return 1 / result
	}

	return result
}

func Test_0050_Example1(t *testing.T) {
	result := myPow(2, 10)
	if result != math.Pow(2, 10) {
		t.Fatalf(fmt.Sprintf("%v", result))
	}
}

func Test_0050_Example2(t *testing.T) {
	result := myPow(2.1, 3)
	if result != math.Pow(2.1, 3) {
		t.Fatalf(fmt.Sprintf("%v", result))
	}
}

func Test_0050_Example3(t *testing.T) {
	result := myPow(2, -2)
	if result != math.Pow(2, -2) {
		t.Fatalf(fmt.Sprintf("%v", result))
	}
}

func Test_0050_Example4(t *testing.T) {
	result := myPow(123, 0)
	if result != math.Pow(123, 0) {
		t.Fatalf(fmt.Sprintf("%v", result))
	}
}
