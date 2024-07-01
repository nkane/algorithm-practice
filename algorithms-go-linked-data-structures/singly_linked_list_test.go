package data_structures

import (
	"testing"

	"gotest.tools/assert"
)

func TestCells(t *testing.T) {
	aCell := Cell[int]{
		Data: 1,
		Next: nil,
	}
	bCell := Cell[int]{
		Data: 2,
		Next: nil,
	}
	aCell.AddAfter(&bCell)

	assert.Assert(t, aCell.Next == &bCell)
	assert.Assert(t, aCell.Data == 1)
}

func TestStack(t *testing.T) {
	stack := CreateLinkedList[int]()
	list := []int{
		4, 2, 0, 6, 9,
	}
	stack.AddRange(list)
	t.Log(stack.ToString(" "))

	for !stack.IsEmpty() {
		t.Logf("Popped: %+v, remaining %d : %s\n",
			*stack.Pop(),
			stack.Length(),
			stack.ToString(" "))
	}
}
