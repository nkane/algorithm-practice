package data_structures

import (
	"fmt"
	"strings"
)

type Cell[T any] struct {
	Data T
	Next *Cell[T]
}

func (c *Cell[T]) AddAfter(after *Cell[T]) {
	after.Next = c.Next
	c.Next = after
}

func (c *Cell[T]) DeleteAfter() *Cell[T] {
	if c.Next == nil {
		panic("expected cell to have next, doesn't have next pointer")
	}
	after := c.Next
	c.Next = after.Next
	return after
}

type LinkedList[T any] struct {
	Sentinel *Cell[T]
}

func CreateLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		Sentinel: &Cell[T]{
			Next: nil,
		},
	}
}

func (list *LinkedList[T]) AddRange(values []T) {
	lastCell := list.Sentinel
	for lastCell.Next != nil {
		lastCell = lastCell.Next
	}
	for _, v := range values {
		lastCell.AddAfter(&Cell[T]{
			Data: v,
			Next: nil,
		})
		lastCell = lastCell.Next
	}
}

func (list *LinkedList[T]) Length() int {
	cell := list.Sentinel.Next
	l := 0
	for cell != nil {
		l++
		cell = cell.Next
	}
	return l
}

func (list *LinkedList[T]) IsEmpty() bool {
	return list.Sentinel.Next == nil
}

func (list *LinkedList[T]) ToString(separator string) string {
	sb := strings.Builder{}
	cell := list.Sentinel.Next
	for cell != nil {
		sb.WriteString(fmt.Sprintf("%+v%s", cell.Data, separator))
		cell = cell.Next
	}
	return sb.String()
}

func (list *LinkedList[T]) ToStringMax(separator string, max int) string {
	sb := strings.Builder{}
	cell := list.Sentinel.Next
	i := 0
	for cell != nil && i < max {
		sb.WriteString(fmt.Sprintf("%+v%s", cell.Data, separator))
		cell = cell.Next
		i++
	}
	return sb.String()
}

// Contains returns true if a particular value is present in the list
func (list *LinkedList[T]) Contains(value T) bool {
	// TODO(nick): optional implementatioen
	return false
}

// Find returns the cell that contains a particular value
// (or possibly the cell before it so you can remove the cell if desired)
func (list *LinkedList[T]) Find(value T) *Cell[T] {
	// TODO(nick): optional implementatioen
	return nil
}

// Remove removes the cell containing a particular value
func (list *LinkedList[T]) Remove(value T) {
	// TODO(nick): optional implementatioen
}

// RemoveAt removes the cell at a particular position
// (such as the 7th cell in the list)
func (list *LinkedList[T]) RemoveAt(index int) {
	// TODO(nick): optional implementatioen
}

// Append adds a new value to the end of the list
func (list *LinkedList[T]) Append(value T) {
	// TODO(nick): optional implementatioen
}

// AddList adds the cells in one list to the end of another,
// like add range to add a list instead of a slice
func (list *LinkedList[T]) Addlist(l *LinkedList[T]) {
	// TODO(nick): optional implementatioen
}

// ToSlice returns a slice containing the list's values
func (list *LinkedList[T]) ToSlice() []T {
	// TODO(nick): optional implementatioen
	return nil
}

// Clone returns a new linked list containing the same cells
func (list *LinkedList[T]) Clone() *LinkedList[T] {
	// TODO(nick): optional implementatioen
	return nil
}

// Clear removes all of the items from the list
func (list *LinkedList[T]) Clear() {
	// TODO(nick): optional implementatioen
}

func (list *LinkedList[T]) Push(value T) {
	list.Sentinel.AddAfter(&Cell[T]{
		Data: value,
		Next: nil,
	})
}

func (list *LinkedList[T]) Pop() *T {
	pop := list.Sentinel.DeleteAfter()
	return &pop.Data
}
