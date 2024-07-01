package data_structures

import (
	"testing"

	"gotest.tools/assert"
)

func TestLoops(t *testing.T) {
	list := CreateLinkedList[string]()
	values := []string{
		"a", "b", "c", "d", "e", "f",
	}
	list.AddRange(values)

	t.Log(list.ToStringMax(" | ", 10))
	value := HasLoop(list)
	assert.Assert(t, value == false, "expected value to be false")

	list.Sentinel.Next.Next.Next.Next.Next.Next = list.Sentinel.Next.Next
	t.Log(list.ToStringMax(" | ", 10))
	value = HasLoop(list)
	assert.Assert(t, value == true, "expected value to be true")
}
