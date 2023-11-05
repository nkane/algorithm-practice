package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicInsertionSort(t *testing.T) {
	testInput := []int{128, 4, 32, 64, 1024, 256, 16}
	testOutput := InsertionSort[int](testInput)
	assert.Equal(t, []int{4, 16, 32, 64, 128, 256, 1024}, testOutput)
}
