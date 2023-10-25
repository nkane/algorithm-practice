package queue

import (
	"week-2/list"
)

// Deque, double-ended queue
type Deque struct {
	list list.SinglyLinkedList
}

func CreateQueue() *Deque {
	queue := Deque{}
	return &queue
}

func (q *Deque) Enqueue(v interface{}) {
	// TODO(nick): implement
}

func (q *Deque) Dequeue() interface{} {
	// TODO(nick): implement
	return nil
}

func (q *Deque) AddFirst(v interface{}) {
	// TODO(nick): implement
}

func (q *Deque) AddLast(v interface{}) {
	// TODO(nick): implement
}

func (q *Deque) RemoveFirst() {
	// TODO(nick): implement
}

func (q *Deque) RemoveLast() {
	// TODO(nick): implement
}

func (q *Deque) Next() interface{} {
	// TODO(nick): implement
	return nil
}

func (q *Deque) Size() int {
	return q.list.Length
}
