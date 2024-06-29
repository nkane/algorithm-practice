package algorithms

import (
	"testing"

	"gotest.tools/assert"
)

func TestQuickSortSlow(t *testing.T) {
	array := []int{
		10, 5, 2, 3,
	}
	expected := []int{
		2, 3, 5, 10,
	}
	got := QuickSortSlow(array)
	assert.DeepEqual(t, expected, got)

	array = MakeRandomSlice(10, 100)
	got = QuickSortSlow(array)
	assert.Assert(t, CheckSort(got) == true)
}

func TestQuickSort(t *testing.T) {
	array := []int{
		10, 5, 2, 3,
	}
	expected := []int{
		2, 3, 5, 10,
	}
	QuickSort(array, 0, len(array)-1)
	assert.DeepEqual(t, expected, array)

	array = MakeRandomSlice(100, 1000)
	QuickSort(array, 0, len(array)-1)
	assert.Assert(t, CheckSort(array) == true)
}
