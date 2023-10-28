package queue

import (
	"week-2/list"
)

// Deque, double-ended queue, first in first out structure
type Deque struct {
	List list.SinglyLinkedList
}

func CreateDeque() *Deque {
	queue := Deque{}
	return &queue
}

func (q *Deque) Enqueue(v interface{}) {
	q.List.Append(v)
}

func (q *Deque) Dequeue() interface{} {
	node := q.List.Find(0)
	if node == nil {
		return nil
	}
	value := node.Value
	q.List.RemoveAt(0)
	return value
}

func (q *Deque) PeekFirst() interface{} {
	return q.List.Head.Value
}

func (q *Deque) PeekLast() interface{} {
	return q.List.Tail.Value
}

func (q *Deque) AddFirst(v interface{}) {
	q.List.Prepend(v)
}

func (q *Deque) AddLast(v interface{}) {
	q.Enqueue(v)
}

func (q *Deque) RemoveFirst() interface{} {
	return q.Dequeue()
}

func (q *Deque) RemoveLast() interface{} {
	node := q.List.Find(q.List.Length - 1)
	if node == nil {
		return nil
	}
	value := node.Value
	q.List.RemoveAt(q.List.Length - 1)
	return value
}
