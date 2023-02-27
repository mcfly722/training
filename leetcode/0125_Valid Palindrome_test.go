package main_test

import (
	"testing"
)

/*
125. Valid Palindrome
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.



Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.


Constraints:

1 <= s.length <= 2 * 105
s consists only of printable ASCII characters.

*/

func isPalindromeString(s string) bool {
	filtered := []rune{}

	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			filtered = append(filtered, c)
		}

		if c >= 'A' && c <= 'Z' {
			filtered = append(filtered, c-'A'+'a')
		}

		if c >= '0' && c <= '9' {
			filtered = append(filtered, c)
		}
	}

	for i := 0; i < len(filtered)/2; i++ {
		if filtered[i] != filtered[len(filtered)-1-i] {
			return false
		}
	}

	//	fmt.Println(string(filtered))

	return true
}

func Test_0125_Example1(t *testing.T) {
	if !isPalindromeString("A man, a plan, a canal: Panama") {
		t.Fatal()
	}
}

func Test_0125_Example2(t *testing.T) {
	if isPalindromeString("race a car") {
		t.Fatal()
	}
}

func Test_0125_Example3(t *testing.T) {
	if !isPalindromeString(" ") {
		t.Fatal()
	}
}

func Test_0125_Example4(t *testing.T) {
	if isPalindromeString("0P") {
		t.Fatal()
	}
}
