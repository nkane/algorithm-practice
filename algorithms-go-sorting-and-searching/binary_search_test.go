package algorithms

import (
	"testing"

	"gotest.tools/assert"
)

func TestBinarySearch(t *testing.T) {
	a := []int{
		1, 2, 3, 4, 5, 6, 7, 8,
	}

	got, testCount := BinarySearchIterative(a, 5)
	assert.Assert(t, got == 4)
	assert.Assert(t, testCount == 3)

	got, testCount = BinarySearchIterative(a, 3)
	assert.Assert(t, got == 2)
	assert.Assert(t, testCount == 3)

	got, testCount = BinarySearchIterative(a, 1)
	assert.Assert(t, got == 0)
	assert.Assert(t, testCount == 3)

	got, testCount = BinarySearchIterative(a, 4)
	assert.Assert(t, testCount == 1)
	assert.Assert(t, got == 3)

	got, testCount = BinarySearchIterative(a, 2)
	assert.Assert(t, testCount == 2)
	assert.Assert(t, got == 1)
}
