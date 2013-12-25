package stack

import (
	"errors"
)

type Stack []interface {
}

func (s *Stack) Push(i interface{}) {
	*s = append(*s, i)
}

func (s Stack) Pushf(i interface{}) Stack {
	return append(s, i)
}

func (s *Stack) Pop() (interface{}, error) {
	stack := *s

	if stack.IsEmpty() {
		return nil, errors.New("Empty Stack!")
	}

	result := stack[stack.Len()-1]
	*s = stack[:stack.Len()-1]

	return result, nil
}

func (s Stack) Popf() (Stack, interface{}, error) {
	if s.IsEmpty() {
		return nil, nil, errors.New("Empty Stack!")
	}

	return s[:s.Len()-1], s[s.Len()-1], nil
}

func (s Stack) Len() int {
	return len(s)
}

func (s Stack) Cap() int {
	return cap(s)
}

func (s Stack) IsEmpty() bool {
	return len(s) < 1
}
