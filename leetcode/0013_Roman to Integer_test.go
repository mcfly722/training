package leetcode_test

import (
	"testing"
)

/*
13. Roman to Integer

Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given a roman numeral, convert it to an integer.



Example 1:

Input: s = "III"
Output: 3
Explanation: III = 3.
Example 2:

Input: s = "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.
Example 3:

Input: s = "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.


Constraints:

1 <= s.length <= 15
s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
It is guaranteed that s is a valid roman numeral in the range [1, 3999].

*/

func romanToInt(s string) int {
	bytes := []byte(s)

	symbols := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	result := 0

	for i := 0; i < len(s); i++ {

		s1 := bytes[i]
		s2 := byte('-')

		if i < len(s)-1 {
			s2 = bytes[i+1]

		}

		if (s1 == 'I' && (s2 == 'V' || s2 == 'X')) || (s1 == 'X' && (s2 == 'L' || s2 == 'C')) || (s1 == 'C' && (s2 == 'D' || s2 == 'M')) {
			result += symbols[s2] - symbols[s1]
			i++
		} else {
			result += symbols[s1]
		}
	}

	return result
}

func Test_0013_Example1(t *testing.T) {
	if !(romanToInt("III") == 3) {
		t.Fatal()
	}
}

func Test_0013_Example2(t *testing.T) {
	if !(romanToInt("LVIII") == 58) {
		t.Fatal()
	}
}

func Test_0013_Example3(t *testing.T) {
	if !(romanToInt("MCMXCIV") == 1994) {
		t.Fatal()
	}
}
