package leetcode_test

import (
	"fmt"
	"sort"
	"testing"
)

/*
49. Group Anagrams
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.



Example 1:

Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
Example 2:

Input: strs = [""]
Output: [[""]]
Example 3:

Input: strs = ["a"]
Output: [["a"]]


Constraints:

1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] consists of lowercase English letters.

*/

func sortLettersInString(str string) string {
	bytes := []byte(str)
	sort.Slice(bytes, func(i int, j int) bool { return bytes[i] < bytes[j] })
	return string(bytes)
}

func groupAnagrams(strs []string) [][]string {
	anagrams := map[string]([]string){}

	for _, str := range strs {

		key := sortLettersInString(str)

		if _, found := anagrams[key]; found {
			anagrams[key] = append(anagrams[key], str)
		} else {
			anagrams[key] = []string{str}
		}
	}

	result := [][]string{}

	for _, val := range anagrams {
		result = append(result, val)
	}

	//fmt.Println(fmt.Sprintf("%v", anagrams))
	return result
}

func Test_0049_Example1(t *testing.T) {
	result := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	if fmt.Sprintf("%v", result) != "[[eat tea ate] [tan nat] [bat]]" {
		t.Fatal(result)
	}
}

func Test_0049_Example2(t *testing.T) {
	result := groupAnagrams([]string{""})

	if fmt.Sprintf("%v", result) != "[[]]" {
		t.Fatal(result)
	}
}

func Test_0049_Example3(t *testing.T) {
	result := groupAnagrams([]string{"a"})

	if fmt.Sprintf("%v", result) != "[[a]]" {
		t.Fatal(result)
	}
}
