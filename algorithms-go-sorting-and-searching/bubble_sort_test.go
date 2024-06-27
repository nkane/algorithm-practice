package algorithms

import (
	"testing"

	"gotest.tools/assert"
)

func TestBubbleSort(t *testing.T) {
	slice := MakeRandomSlice(10, 11)
	PrintSlice(slice, 10)
	BubbleSort(slice)
	PrintSlice(slice, 10)
	got := CheckSort(slice)
	assert.Assert(t, true == got, "Expected: %+v, Got: %+v\n", true, got)
}

func TestBubbleSortOptimized(t *testing.T) {
	slice := MakeRandomSlice(10, 11)
	PrintSlice(slice, 10)
	BubbleSortOptimized(slice)
	PrintSlice(slice, 10)
	got := CheckSort(slice)
	assert.Assert(t, true == got, "Expected: %+v, Got: %+v\n", true, got)
}

func TestCocktailShakerSort(t *testing.T) {
	slice := MakeRandomSlice(10, 11)
	PrintSlice(slice, 10)
	BubbleSortOptimized(slice)
	PrintSlice(slice, 10)
	got := CheckSort(slice)
	assert.Assert(t, true == got, "Expected: %+v, Got: %+v\n", true, got)
}
