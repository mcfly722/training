package leetcode_test

import (
	"fmt"
	"testing"
)

/*
14. Longest Common Prefix
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".



Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.


Constraints:

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lowercase English letters.

*/

func longestCommonPrefix(strs []string) string {
	prefix := ""

	for i := 0; ; i++ {

		if i > len(strs[0])-1 {
			return prefix
		} else {
			c := (strs[0])[i]

			for _, str := range strs {
				if i > len(str)-1 || str[i] != c {
					return prefix
				}
			}

			prefix += string(c)

		}

	}
}

func Test_0014_Example1(t *testing.T) {
	prefix := longestCommonPrefix([]string{"flower", "flow", "flight"})
	if prefix != "fl" {
		t.Fatalf(fmt.Sprintf("prefix:%v", prefix))
	}
}

func Test_0014_Example2(t *testing.T) {
	prefix := longestCommonPrefix([]string{"dog", "racecar", "car"})
	if prefix != "" {
		t.Fatalf(fmt.Sprintf("prefix:%v", prefix))
	}
}

func Test_0014_Example3(t *testing.T) {
	prefix := longestCommonPrefix([]string{"do", "doing", "door"})
	if prefix != "do" {
		t.Fatalf(fmt.Sprintf("prefix:%v", prefix))
	}
}
