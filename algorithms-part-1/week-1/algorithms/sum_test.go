package algorithms

import "testing"

func TestTwoSumSlow(t *testing.T) {
	testInput := []int{0, 1, 2, -2, 3, 1}
	expectedOutput := 1
	output := TwoSumSlow(testInput)
	if output != expectedOutput {
		t.Fatalf("TwoSumSlow output does not match expected output")
	}
}

func TestTwoSumFast(t *testing.T) {
	testInput := []int{-2, 0, 1, 1, 2, 3}
	expectedOutput := 1
	output := TwoSumFast(testInput)
	if output != expectedOutput {
		t.Fatalf("TwoSumSlow output does not match expected output")
	}
}

func TestThreeSumSlow(t *testing.T) {
	testInput := []int{0, 1, -1, -10, 10}
	expectedOutput := 2
	output := ThreeSumSlow(testInput)
	if output != expectedOutput {
		t.Fatalf("ThreeSumSlow output does not match expected output %d", output)
	}
}

func TestThreeSumFast(t *testing.T) {
	testInput := []int{-10, -1, 0, 1, 10}
	expectedOutput := 2
	output := ThreeSumFast(testInput)
	if output != expectedOutput {
		t.Fatalf("ThreeSumFast output does not match expected output %d", output)
	}
}
