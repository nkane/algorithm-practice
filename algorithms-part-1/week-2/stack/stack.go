package stack

import (
	"week-2/list"
)

// Stack, first in last out structure
type Stack struct {
	List list.SinglyLinkedList
}

func CreateStack() *Stack {
	stack := Stack{}
	return &stack
}

func (s *Stack) Push(v interface{}) {
	s.List.Prepend(v)
}

func (s *Stack) Pop() interface{} {
	node := s.List.Find(0)
	if node == nil {
		return nil
	}
	value := node.Value
	s.List.RemoveAt(0)
	return value
}

func (s *Stack) Peek() interface{} {
	return s.List.Find(0)
}
