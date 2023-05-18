package leetcode_test

import (
	"testing"
)

/*
Write a function that reverses a string. The input string is given as an array of characters s.

You must do this by modifying the input array in-place with O(1) extra memory.



Example 1:

Input: s = ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]
Example 2:

Input: s = ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]


Constraints:

1 <= s.length <= 105
s[i] is a printable ascii character.
   Hide Hint #1
The entire logic for reversing a string is based on using the opposite directional two-pointer approach!
*/

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		tmp1 := s[i]
		s[i] = s[len(s)-1-i]
		s[len(s)-1-i] = tmp1
	}
}

func Test_1440_Example1(t *testing.T) {
	s := []byte("hello")
	reverseString(s)
	if !(string(s) == "olleh") {
		t.Fatal(string(s))
	}
}

func Test_1440_Example2(t *testing.T) {
	s := []byte("Hannah")
	reverseString(s)
	if !(string(s) == "hannaH") {
		t.Fatal(string(s))
	}
}
