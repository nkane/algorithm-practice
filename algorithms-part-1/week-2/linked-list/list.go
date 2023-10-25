package list

import "errors"

type Node struct {
	Next  *Node
	Value interface{}
}

type SingleLinkedList struct {
	Head        *Node
	Tail        *Node
	Length      int
	IteratorIdx int
}

func CreateSingleLinkedList() *SingleLinkedList {
	list := SingleLinkedList{}
	return &list
}

func (l *SingleLinkedList) Append(v interface{}) {
	n := Node{
		Next:  nil,
		Value: v,
	}
	l.Length++
	if l.Head == nil {
		l.Head = &n
		l.Tail = &n
		return
	}
	l.Tail.Next = &n
	l.Tail = &n
}

func (l *SingleLinkedList) Prepend(v interface{}) {
	n := Node{
		Next:  nil,
		Value: v,
	}
	l.Length++
	if l.Head == nil {
		l.Head = &n
		l.Tail = &n
		return
	}
	n.Next = l.Head
	l.Head = &n
}

func (l *SingleLinkedList) Find(idx int) *Node {
	node := l.Head
	sIdx := 0
	for node != nil && sIdx < idx {
		node = node.Next
		sIdx++
	}
	return node
}

func (l *SingleLinkedList) RemoveAt(idx int) error {
	if idx > l.Length-1 {
		return errors.New("cannot remove, index out of range")
	}
	if idx == 0 {
		prevHead := l.Head
		l.Head = l.Head.Next
		prevHead.Next = nil
		prevHead.Value = nil
		l.Length--
		if l.Length == 0 {
			l.Tail = nil
		}
		return nil
	}
	prevNode := l.Find(idx - 1)
	if idx == l.Length-1 {
		prevTail := l.Tail
		l.Tail = prevNode
		prevTail.Next = nil
		prevTail.Value = nil
		l.Length--
		return nil
	}
	node := l.Find(idx)
	prevNode.Next = node.Next
	node.Next = nil
	node.Value = nil
	l.Length--
	return nil
}

func (l *SingleLinkedList) Next() *Node {
	if l.Length < l.IteratorIdx {
		return nil
	}
	node := l.Find(l.IteratorIdx)
	if node == nil {
		return nil
	}
	l.IteratorIdx++
	return node
}
