package data_structures

import (
	"fmt"
	"strings"
)

type Node[T any] struct {
	Data     T
	Next     *Node[T]
	Previous *Node[T]
}

func (n *Node[T]) AddAfter(after *Node[T]) {
	after.Previous = n
	after.Next = n.Next
	n.Next.Previous = after
	n.Next = after
}

func (n *Node[T]) AddBefore(before *Node[T]) {
	n.Previous.AddAfter(before)
}

func (n *Node[T]) DeleteAfter() *Node[T] {
	//
	if n.Next == nil {
		panic("expected node to have next, doesn't have next pointer")
	}
	after := n.Next
	n.Next = after.Next
	n.Next.Previous = n
	return after
}

type DoublyLinkedList[T any] struct {
	TopSentinel    *Node[T]
	BottomSentinel *Node[T]
}

func CreateDoublyLinkedList[T any](topData T, bottomData T) *DoublyLinkedList[T] {
	l := &DoublyLinkedList[T]{
		TopSentinel: &Node[T]{
			Previous: nil,
			Data:     topData,
		},
		BottomSentinel: &Node[T]{
			Next: nil,
			Data: bottomData,
		},
	}
	l.TopSentinel.Next = l.BottomSentinel
	l.BottomSentinel.Previous = l.TopSentinel
	return l
}

func (list *DoublyLinkedList[T]) AddRange(values []T) {
	lastNode := list.BottomSentinel
	for _, v := range values {
		lastNode.AddBefore(&Node[T]{
			Data:     v,
			Next:     nil,
			Previous: nil,
		})
		// lastNode = lastNode.Next
	}
}

func (list *DoublyLinkedList[T]) Length() int {
	node := list.TopSentinel.Next
	l := 0
	for node != list.BottomSentinel {
		l++
		node = node.Next
	}
	return l
}

func (list *DoublyLinkedList[T]) IsEmpty() bool {
	return list.TopSentinel.Next == nil
}

func (list *DoublyLinkedList[T]) ToString(separator string) string {
	sb := strings.Builder{}
	n := list.TopSentinel.Next
	for n != list.BottomSentinel {
		sb.WriteString(fmt.Sprintf("%+v%s", n.Data, separator))
		n = n.Next
	}
	return sb.String()
}

func (list *DoublyLinkedList[T]) ToStringMax(separator string, max int) string {
	sb := strings.Builder{}
	n := list.TopSentinel.Next
	i := 0
	for n != list.BottomSentinel && i < max {
		sb.WriteString(fmt.Sprintf("%+v%s", n.Data, separator))
		n = n.Next
		i++
	}
	return sb.String()
}

/*
func (list *DoublyLinkedList[T]) Push(value T) {
	list.Sentinel.AddAfter(&Cell[T]{
		Data: value,
		Next: nil,
	})
}

func (list *DoublyLinkedList[T]) Pop() *T {
	pop := list.Sentinel.DeleteAfter()
	return &pop.Data
}
*/
