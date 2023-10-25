package stack

import (
	"week-2/list"
)

type Stack struct {
	list list.SinglyLinkedList
}

func CreateStack() *Stack {
	stack := Stack{}
	return &stack
}

func (s *Stack) Push(v interface{}) {
	// TODO(nick): implement
}

func (s *Stack) Pop() interface{} {
	// TODO(nick): implement
	return nil
}

func (s *Stack) Peek() interface{} {
	// TODO(nick): implement
	return nil
}
