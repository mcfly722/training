package leetcode_test

import "testing"

/*
242. Valid Anagram
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.



Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false


Constraints:

1 <= s.length, t.length <= 5 * 104
s and t consist of lowercase English letters.


Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?
*/

func isAnagram(s string, t string) bool {
	chars := map[rune]int{}

	for _, c := range s {
		chars[c]++
	}

	for _, c := range t {
		if _, found := chars[c]; found {
			if chars[c] == 1 {
				delete(chars, c)
			} else {
				chars[c]--
			}
		} else {
			return false
		}
	}

	return len(chars) == 0
}

func Test_0242_Example1(t *testing.T) {
	if !isAnagram("anagram", "nagaram") {
		t.Fatal()
	}
}

func Test_0242_Example2(t *testing.T) {
	if isAnagram("rat", "car") {
		t.Fatal()
	}
}
