package main_test

import (
	"testing"
)

/*
383. Ransom Note
Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.

Each letter in magazine can only be used once in ransomNote.



Example 1:

Input: ransomNote = "a", magazine = "b"
Output: false
Example 2:

Input: ransomNote = "aa", magazine = "ab"
Output: false
Example 3:

Input: ransomNote = "aa", magazine = "aab"
Output: true


Constraints:

1 <= ransomNote.length, magazine.length <= 105
ransomNote and magazine consist of lowercase English letters.
*/

func canConstruct(ransomNote string, magazine string) bool {
	letters := map[rune]int{}

	for _, letter := range magazine {
		letters[letter]++
	}

	for _, letter := range ransomNote {
		if count, found := letters[letter]; !found || count == 0 {
			return false
		}
		letters[letter]--
	}

	return true
}

func Test_0383_Example1(t *testing.T) {
	if canConstruct("a", "b") != false {
		t.Fatal()
	}
}

func Test_0383_Example2(t *testing.T) {
	if canConstruct("aa", "ab") != false {
		t.Fatal()
	}
}

func Test_0383_Example3(t *testing.T) {
	if canConstruct("aa", "aab") != true {
		t.Fatal()
	}
}
