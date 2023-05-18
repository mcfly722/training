package leetcode_test

import "errors"

type Stack []byte

func (s *Stack) Push(c byte) {
	*s = append(*s, c)
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Pop() (byte, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, nil
	}
}
