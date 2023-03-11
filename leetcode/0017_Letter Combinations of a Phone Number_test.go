package main_test

import (
	"fmt"
	"testing"
)

/*
17. Letter Combinations of a Phone Number
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.




Example 1:

Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
Example 2:

Input: digits = ""
Output: []
Example 3:

Input: digits = "2"
Output: ["a","b","c"]


Constraints:

0 <= digits.length <= 4
digits[i] is a digit in the range ['2', '9'].

*/

func recursiveLetterCombinations(digits string, currentPath string, onFound func(combination string)) {
	alphabet := map[int]([]rune){
		2: []rune{'a', 'b', 'c'},
		3: []rune{'d', 'e', 'f'},
		4: []rune{'g', 'h', 'i'},
		5: []rune{'j', 'k', 'l'},
		6: []rune{'m', 'n', 'o'},
		7: []rune{'p', 'q', 'r', 's'},
		8: []rune{'t', 'u', 'v'},
		9: []rune{'w', 'x', 'y', 'z'},
	}

	if len(digits) == 0 {
		onFound(currentPath)
	} else {
		if letters, found := alphabet[int(digits[0]-'0')]; found {
			for _, letter := range letters {
				recursiveLetterCombinations(digits[1:], currentPath+string(letter), onFound)
			}
		}
	}

}

func letterCombinations(digits string) []string {
	solutions := []string{}

	if digits == "" {
		return solutions
	}

	onFound := func(combination string) {
		solutions = append(solutions, combination)
	}

	recursiveLetterCombinations(digits, "", onFound)

	return solutions
}

func Test_0017_Example1(t *testing.T) {
	result := fmt.Sprintf("%v", letterCombinations("23"))
	if result != "[ad ae af bd be bf cd ce cf]" {
		t.Fatal(result)
	}
}

func Test_0017_Example2(t *testing.T) {
	result := letterCombinations("")
	if len(result) != 0 {
		t.Fatal(fmt.Sprintf("%v", result))
	}
}
