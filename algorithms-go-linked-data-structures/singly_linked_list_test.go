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
