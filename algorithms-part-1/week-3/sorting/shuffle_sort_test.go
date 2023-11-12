package sorting

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestShuffleSort(t *testing.T) {
	seed := int64(256)
	list := []int{0, 1, 2, 3, 4, 5}
	expectedShuffleString := fmt.Sprintf("%+v", []int{0, 3, 2, 5, 4, 1})
	ShuffleSort[int](list, seed)
	listString := fmt.Sprintf("%+v", list)
	assert.Equal(t, expectedShuffleString, listString)
}
