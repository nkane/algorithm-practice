package binary_search

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestBinarySerach(t *testing.T) {
	list := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
	result := BinarySearch(list, 6)
	expectedIndex := 5
	assert.Assert(t, expectedIndex == result, "expected: %d, got:%d", expectedIndex, result)
}
