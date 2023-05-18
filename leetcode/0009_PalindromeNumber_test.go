package leetcode_test

import (
	"fmt"
	"testing"
)

/*
9. Palindrome Number

Given an integer x, return true if x is a
palindrome
, and false otherwise.



Example 1:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.
Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.


Constraints:

-231 <= x <= 231 - 1


Follow up: Could you solve it without converting the integer to a string?
*/

func num2arrReversed(x int) []byte {
	if x == 0 {
		return []byte{0}
	}
	result := []byte{}

	for ; x > 0; x = x / 10 {
		result = append(result, byte(x%10))
	}

	return result
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	bytesRepresentation := num2arrReversed(x)
	for i := 0; i < len(bytesRepresentation)/2; i++ {
		if bytesRepresentation[i] != bytesRepresentation[len(bytesRepresentation)-1-i] {
			return false
		}
	}

	return true
}

func Test_0009_Example4(t *testing.T) {
	result := num2arrReversed(0)
	if !IsEqualByteArrays(result, []byte{0}) {
		t.Fatalf(fmt.Sprintf("%v", result))
	}
}

func Test_0009_Example5(t *testing.T) {
	result := num2arrReversed(12345670)
	if !IsEqualByteArrays(result, []byte{0, 7, 6, 5, 4, 3, 2, 1}) {
		t.Fatalf(fmt.Sprintf("%v", result))
	}
}

func Test_0009_Example1(t *testing.T) {
	if !isPalindrome(121) {
		t.Fatal()
	}
}

func Test_0009_Example2(t *testing.T) {
	if isPalindrome(-121) {
		t.Fatal()
	}
}

func Test_0009_Example3(t *testing.T) {
	if isPalindrome(10) {
		t.Fatal()
	}
}
