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

func (n *Node[T]) Delete() *Node[T] {
	if n.Next == nil {
		panic("expected node to have next, doesn't have next pointer")
	}
	next := n.Next
	previous := n.Previous
	next.Previous = previous
	previous.Next = next
	return n
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
	return list.TopSentinel.Next == list.BottomSentinel
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

func (list *DoublyLinkedList[T]) Enqueue(value T) {
	list.TopSentinel.AddAfter(&Node[T]{
		Data: value,
	})
}

func (list *DoublyLinkedList[T]) PushTop(value T) {
	list.TopSentinel.AddAfter(&Node[T]{
		Data: value,
	})
}

func (list *DoublyLinkedList[T]) PopTop() *Node[T] {
	if list.IsEmpty() {
		return nil
	}
	return list.TopSentinel.Next.Delete()
}

func (list *DoublyLinkedList[T]) PushBottom(value T) {
	list.BottomSentinel.AddBefore(&Node[T]{
		Data: value,
	})
}

func (list *DoublyLinkedList[T]) PopBottom() *Node[T] {
	if list.IsEmpty() {
		return nil
	}
	return list.BottomSentinel.Previous.Delete()
}

func (list *DoublyLinkedList[T]) Dequeue() *Node[T] {
	return list.BottomSentinel.Previous.Delete()
}

func (list *DoublyLinkedList[T]) Deque() *Node[T] {
	return list.PopBottom()
}
