package data_structures

import (
	"errors"
	"log"
)

type Node struct {
	Next  *Node       `json:"next"`
	Value interface{} `json:"value"`
}

type SinglyLinkedList struct {
	Head   *Node `json:"head"`
	Tail   *Node `json:"tail"`
	Length int   `json:"length"`
}

func (l *SinglyLinkedList) Prepend(v interface{}) (*Node, error) {
	if l.Length == 0 {
		return l.Append(v)
	}
	l.Length++
	n := &Node{
		Next:  l.Head,
		Value: v,
	}
	l.Head = n
	return n, nil
}

func (l *SinglyLinkedList) Append(v interface{}) (*Node, error) {
	n := &Node{
		Next:  nil,
		Value: v,
	}
	if l.Length == 0 {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = n
	}
	l.Length++
	return n, nil
}

func (l *SinglyLinkedList) Insert(v interface{}, idx int) (*Node, error) {
	if l.Length == 0 || idx == l.Length-1 {
		return l.Append(v)
	}
	if idx == 0 {
		return l.Prepend(v)
	}
	// TODO(nick): finish this
	return nil, nil
}

func (l *SinglyLinkedList) Remove(idx int) interface{} {
	return nil
}

func (l *SinglyLinkedList) Get(idx int) (*Node, error) {
	if (idx > (l.Length - 1)) || (idx < 0) {
		return nil, errors.New("index out of range")
	}
	n := l.Head
	for i := 0; i < idx; i++ {
		n = n.Next
		if n == nil {
			return nil, errors.New("node is nil")
		}
	}
	return n, nil
}

func (l *SinglyLinkedList) DumpInfo() {
	n := l.Head
	log.Printf("List Length: %d\n", l.Length)
	for n != nil {
		log.Printf("%p - %+v\n", n, n)
		n = n.Next
	}
}
