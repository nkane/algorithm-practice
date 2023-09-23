package algorithms

import "testing"

func TestBinarySearchExpectFound(t *testing.T) {
	testInput := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	expectedOutput := 4
	output := BinarySearch(4, testInput)
	if output != expectedOutput {
		t.Fatalf("BinarySearch output does not match expected output")
	}
}

func TestBinarySearchExpectNotFound(t *testing.T) {
	testInput := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	expectedOutput := -1
	output := BinarySearch(10, testInput)
	if output != expectedOutput {
		t.Fatalf("BinarySearch output does not match expected output")
	}
}
