package main_test

import (
	"errors"
	"testing"
)

/*
20. Valid Parentheses
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.


Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false


Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/

type Stack []byte

func (s *Stack) Push(c byte) {
	*s = append(*s, c)
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Pop() (error, byte) {
	if s.IsEmpty() {
		return errors.New("stack is empty"), 0
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return nil, element
	}
}

func isValid(s string) bool {
	closingPairs := map[byte]byte{')': '(', '}': '{', ']': '['}
	var stack Stack

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '{', '(', '[':
			stack.Push(s[i])
		case '}', ')', ']':
			err, c := stack.Pop()
			if err != nil || c != closingPairs[s[i]] {
				return false
			}
		default:
			return false
		}
	}

	err, _ := stack.Pop() // stack finished
	return err != nil
}

func Test_0020_Example1(t *testing.T) {
	if !isValid("()") {
		t.Fatal()
	}
}

func Test_0020_Example2(t *testing.T) {
	if !isValid("()[]{}") {
		t.Fatal()
	}
}

func Test_0020_Example3(t *testing.T) {
	if isValid("(]") {
		t.Fatal()
	}
}

func Test_0020_Example4(t *testing.T) {
	if !isValid("(()[{}]){}[()({[]})]") {
		t.Fatal()
	}
}

func Test_0020_Example5(t *testing.T) {
	if isValid("(") {
		t.Fatal()
	}
}

func Test_0020_Example6(t *testing.T) {
	if !isValid("(((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((())))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))") {
		t.Fatal()
	}
}

func Test_0020_Example7(t *testing.T) {
	if isValid("(((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((())))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))]") {
		t.Fatal()
	}
}

func Test_0020_Example8(t *testing.T) {
	if isValid("(){}}{") {
		t.Fatal()
	}
}
