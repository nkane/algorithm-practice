package algorithms

import (
	"testing"

	"gotest.tools/assert"
)

func TestCountingSort(t *testing.T) {
	numItems := 10
	max := 100
	slice := MakeCustomerSlice(numItems, max)
	shouldBeSorted := CountingSort(slice, 0, max)
	got := CheckCustomersSliceSorted(shouldBeSorted)
	assert.Assert(t, got == true)
}
