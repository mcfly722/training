package leetcode_test

import "testing"

/*
387. First Unique Character in a String
Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.

Example 1:

Input: s = "leetcode"
Output: 0
Example 2:

Input: s = "loveleetcode"
Output: 2
Example 3:

Input: s = "aabb"
Output: -1


Constraints:

1 <= s.length <= 105
s consists of only lowercase English letters.

*/

func firstUniqChar(s string) int {
	chars := map[rune]int{}

	for _, c := range s {
		chars[c]++
	}

	for i, symbol := range s {
		amount, _ := chars[symbol]
		if amount < 2 {
			return i
		}

	}

	return -1
}

func Test_0387_Example1(t *testing.T) {
	if firstUniqChar("leetcode") != 0 {
		t.Fatal()
	}
}

func Test_0387_Example2(t *testing.T) {
	if firstUniqChar("loveleetcode") != 2 {
		t.Fatal()
	}
}

func Test_0387_Example3(t *testing.T) {
	if firstUniqChar("aabb") != -1 {
		t.Fatal()
	}
}
