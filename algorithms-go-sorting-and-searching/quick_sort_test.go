package algorithms

import (
	"testing"

	"gotest.tools/assert"
)

func TestQuickSort(t *testing.T) {
	array := []int{
		10, 5, 2, 3,
	}
	expected := []int{
		2, 3, 5, 10,
	}
	got := QuickSort(array)
	assert.DeepEqual(t, expected, got)
}
