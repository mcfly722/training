package leetcode_test

import (
	"math"
	"testing"
)

/*
3. Longest Substring Without Repeating Characters
Given a string s, find the length of the longest
substring
 without repeating characters.



Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.


Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.

*/

func lengthOfLongestSubstring(s string) int {
	end := 0
	max := 0

	uniques := make([]int, 256)
	chars := []byte(s)

	for start := 0; start < len(chars); {

		if uniques[chars[start]] == 0 {
			uniques[chars[start]] = 1
			start++
			max = (int)(math.Max((float64)(max), (float64)(start-end)))
		} else {
			uniques[chars[end]] = 0
			end++
		}

	}

	return max
}

func Test_0003_Example1(t *testing.T) {
	if lengthOfLongestSubstring("abcabcbb") != 3 {
		t.Fatal()
	}
}

func Test_0003_Example2(t *testing.T) {
	if lengthOfLongestSubstring("bbbbb") != 1 {
		t.Fatal()
	}
}

func Test_0003_Example3(t *testing.T) {
	if lengthOfLongestSubstring("pwwkew") != 3 {
		t.Fatal()
	}
}
