package sorting

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestBasicMergeSort(t *testing.T) {
	testInput := []int{128, 4, 32, 64, 1024, 256, 16}
	testOutput := MergeSort[int](testInput)
	testOutputString := fmt.Sprint(testOutput)

	expectedOutput := []int{4, 16, 32, 64, 128, 256, 1024}
	expectedOutputString := fmt.Sprint(expectedOutput)

	assert.Equal(t, expectedOutputString, testOutputString)
}

func TestMergeSort(t *testing.T) {
	testInput := []int{2048, 128, 4, 32, 64, 1024, 256, 16, -256, 0, 8, 2, -128}
	testOutput := MergeSort[int](testInput)
	testOutputString := fmt.Sprint(testOutput)

	expectedOutput := []int{-256, -128, 0, 2, 4, 8, 16, 32, 64, 128, 256, 1024, 2048}
	expectedOutputString := fmt.Sprint(expectedOutput)

	assert.Equal(t, expectedOutputString, testOutputString)
}
