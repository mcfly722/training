package leetcode_test

import (
	"testing"
)

func longestPalindrome(s string) string {

	longest := ""

	for i := 0; i < len(s); i++ {

		for j := 0; i-j >= 0 && i+j < len(s) && s[i-j] == s[i+j]; j++ {
			if len(longest) < j*2+1 {
				longest = s[i-j : i+j+1]
			}
		}

		for j := 0; i-j >= 0 && i+j+1 < len(s) && s[i-j] == s[i+j+1]; j++ {
			if len(longest) <= j*2+1 {
				longest = s[i-j : i+j+2]
			}
		}

	}

	return longest
}

func Test_0005_Example1(t *testing.T) {
	result := longestPalindrome("babad")
	if result != "bab" {
		t.Fatal(result)
	}
}

func Test_0005_Example2(t *testing.T) {
	result := longestPalindrome("cbbd")
	if result != "bb" {
		t.Fatal(result)
	}
}

func Test_0005_Example3(t *testing.T) {
	result := longestPalindrome("BBac")
	if result != "BB" {
		t.Fatal(result)
	}
}
func Test_0005_Example4(t *testing.T) {
	result := longestPalindrome("acBB")
	if result != "BB" {
		t.Fatal(result)
	}
}

func Test_0005_Example5(t *testing.T) {
	result := longestPalindrome("acBBB")
	if result != "BBB" {
		t.Fatal(result)
	}
}

func Test_0005_Example6(t *testing.T) {
	result := longestPalindrome("BBBac")
	if result != "BBB" {
		t.Fatal(result)
	}
}

func Test_0005_Example7(t *testing.T) {
	result := longestPalindrome("1")
	if result != "1" {
		t.Fatal(result)
	}
}

func Test_0005_Example8(t *testing.T) {
	result := longestPalindrome("")
	if result != "" {
		t.Fatal(result)
	}
}
