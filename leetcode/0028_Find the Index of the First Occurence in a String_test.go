package leetcode_test

import "testing"

/*
28. Find the Index of the First Occurrence in a String
Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.
Example 2:

Input: haystack = "leetcode", needle = "leeto"
Output: -1
Explanation: "leeto" did not occur in "leetcode", so we return -1.


Constraints:

1 <= haystack.length, needle.length <= 104
haystack and needle consist of only lowercase English characters.

*/

func strStr(haystack string, needle string) int {

	for i := 0; i < len(haystack)-len(needle)+1; i++ {

		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
			if j == len(needle)-1 {
				return i
			}
		}

	}

	return -1
}

func Test_0028_Example1(t *testing.T) {
	result := strStr("sadbutsad", "sad")
	if result != 0 {
		t.Fatal(result)
	}
}

func Test_0028_Example2(t *testing.T) {
	result := strStr("leetcode", "leeto")
	if result != -1 {
		t.Fatal(result)
	}
}

func Test_0028_Example3(t *testing.T) {
	result := strStr("eiruthtest", "test")
	if result != 6 {
		t.Fatal(result)
	}
}

func Test_0028_Example4(t *testing.T) {
	result := strStr("", "")
	if result != -1 {
		t.Fatal(result)
	}
}

func Test_0028_Example5(t *testing.T) {
	result := strStr("qwe", "qweqwe")
	if result != -1 {
		t.Fatal(result)
	}
}

func Test_0028_Example6(t *testing.T) {
	result := strStr("qwe", "qwe")
	if result != 0 {
		t.Fatal(result)
	}
}
