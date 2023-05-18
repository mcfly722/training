package leetcode_test

import (
	"fmt"
	"testing"
)

/*
67. Add Binary
Given two binary strings a and b, return their sum as a binary string.



Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"


Constraints:

1 <= a.length, b.length <= 104
a and b consist only of '0' or '1' characters.
Each string does not contain leading zeros except for the zero itself.

*/

func addBinary(a string, b string) string {
	counter := 0

	sum := ""

	for i := 0; i < len(a) || i < len(b); i++ {

		c1 := 0
		c2 := 0

		if i < len(a) && a[len(a)-1-i] == '1' {
			c1 = 1
		}
		if i < len(b) && b[len(b)-1-i] == '1' {
			c2 = 1
		}

		counter = counter + c1 + c2

		sum = fmt.Sprintf("%v%v", counter%2, sum)

		counter = counter >> 1

	}

	for ; counter > 0; counter = counter >> 1 {
		sum = fmt.Sprintf("%v%v", counter%2, sum)
	}

	return sum
}

func Test_0067_Example1(t *testing.T) {
	result := addBinary("11", "1")
	if result != "100" {
		t.Fatal(fmt.Sprintf("%v", result))
	}
}

func Test_0067_Example2(t *testing.T) {
	result := addBinary("1010", "1011")
	if result != "10101" {
		t.Fatal(fmt.Sprintf("%v", result))
	}
}

func Test_0067_Example3(t *testing.T) {
	result := addBinary("1010", "0")
	if result != "1010" {
		t.Fatal(fmt.Sprintf("%v", result))
	}
}

func Test_0067_Example4(t *testing.T) {
	result := addBinary("11111", "1")
	if result != "100000" {
		t.Fatal(fmt.Sprintf("%v", result))
	}
}
