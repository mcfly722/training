package leetcode_test

import "errors"

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
