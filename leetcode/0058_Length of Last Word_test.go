package leetcode_test

import (
	"strings"
	"testing"
)

/*
58. Length of Last Word
Given a string s consisting of words and spaces, return the length of the last word in the string.

A word is a maximal
substring
 consisting of non-space characters only.



Example 1:

Input: s = "Hello World"
Output: 5
Explanation: The last word is "World" with length 5.
Example 2:

Input: s = "   fly me   to   the moon  "
Output: 4
Explanation: The last word is "moon" with length 4.
Example 3:

Input: s = "luffy is still joyboy"
Output: 6
Explanation: The last word is "joyboy" with length 6.


Constraints:

1 <= s.length <= 104
s consists of only English letters and spaces ' '.
There will be at least one word in s.
*/

func lengthOfLastWord(s string) int {

	if s == "" {
		return 0
	}

	strs := strings.Split(s, " ")

	lastLen := 0

	for _, str := range strs {
		if str != "" {
			lastLen = len(str)
		}
	}

	return lastLen
}

func Test_0058_Example1(t *testing.T) {
	if !(lengthOfLastWord("Hello World") == 5) {
		t.Fatal()
	}
}

func Test_0058_Example2(t *testing.T) {
	if !(lengthOfLastWord("   fly me   to   the moon  ") == 4) {
		t.Fatal()
	}
}

func Test_0058_Example3(t *testing.T) {
	if !(lengthOfLastWord("luffy is still joyboy") == 6) {
		t.Fatal()
	}
}
