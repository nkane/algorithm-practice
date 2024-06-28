package algorithms

import (
	"testing"

	"gotest.tools/assert"
)

func TestSum(t *testing.T) {
	expected := 12
	arr := []int{
		2, 4, 6,
	}
	got := Sum(arr)
	assert.Assert(t, expected == got)
}

func TestCount(t *testing.T) {
	expected := 6
	arr := []int{
		2, 4, 6, 8, 10, 12,
	}
	got := Count(arr)
	assert.Assert(t, expected == got)
}

func TestMax(t *testing.T) {
	expected := 22
	arr := []int{
		2, 8, 6, 10, 12, 7, 22,
	}
	got := Max(arr)
	assert.Assert(t, expected == got)
}
