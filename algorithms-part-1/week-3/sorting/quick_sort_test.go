package sorting

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestQuickSort(t *testing.T) {
	testInput := []int{128, 4, 32, 64, 1024, 256, 16}
	testOutput := QuickSort[int](testInput)
	testOutputString := fmt.Sprint(testOutput)

	expectedOutput := []int{4, 16, 32, 64, 128, 256, 1024}
	expectedOutputString := fmt.Sprint(expectedOutput)

	assert.Equal(t, expectedOutputString, testOutputString)
}
