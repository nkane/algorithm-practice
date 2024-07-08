package recursion

import (
	"testing"

	"gotest.tools/assert"
)

func TestFibonacciOnTheFly(t *testing.T) {
	tests := []struct {
		Input  int64
		Output int64
	}{
		{
			Input:  0,
			Output: 0,
		},
		{
			Input:  1,
			Output: 1,
		},
		{
			Input:  2,
			Output: 1,
		},
		{
			Input:  3,
			Output: 2,
		},
		{
			Input:  4,
			Output: 3,
		},
		{
			Input:  5,
			Output: 5,
		},
		{
			Input:  10,
			Output: 55,
		},
		{
			Input:  15,
			Output: 610,
		},
		{
			Input:  20,
			Output: 6765,
		},
		{
			Input:  25,
			Output: 75025,
		},
		{
			Input:  25,
			Output: 75025,
		},
		{
			Input:  30,
			Output: 832040,
		},
		{
			Input:  35,
			Output: 9227465,
		},
		{
			Input:  40,
			Output: 102334155,
		},
	}

	for _, test := range tests {
		out := FibonacciOnTheFly(test.Input)
		assert.Assert(t, out == test.Output, "input: %d, expected: %d, got: %d", test.Input, test.Output, out)
	}
}

func TestFibonacciPrefilled(t *testing.T) {
	tests := []struct {
		Input  int64
		Output int64
	}{
		{
			Input:  0,
			Output: 0,
		},
		{
			Input:  1,
			Output: 1,
		},
		{
			Input:  2,
			Output: 1,
		},
		{
			Input:  3,
			Output: 2,
		},
		{
			Input:  4,
			Output: 3,
		},
		{
			Input:  5,
			Output: 5,
		},
		{
			Input:  10,
			Output: 55,
		},
		{
			Input:  15,
			Output: 610,
		},
		{
			Input:  20,
			Output: 6765,
		},
		{
			Input:  25,
			Output: 75025,
		},
		{
			Input:  25,
			Output: 75025,
		},
		{
			Input:  30,
			Output: 832040,
		},
		{
			Input:  35,
			Output: 9227465,
		},
		{
			Input:  40,
			Output: 102334155,
		},
	}

	InitializeSlice()
	for _, test := range tests {
		out := FibonacciPrefilled(test.Input)
		assert.Assert(t, out == test.Output, "input: %d, expected: %d, got: %d", test.Input, test.Output, out)
	}
}

func TestFibonacciBottomUp(t *testing.T) {
	tests := []struct {
		Input  int64
		Output int64
	}{
		{
			Input:  0,
			Output: 0,
		},
		{
			Input:  1,
			Output: 1,
		},
		{
			Input:  2,
			Output: 1,
		},
		{
			Input:  3,
			Output: 2,
		},
		{
			Input:  4,
			Output: 3,
		},
		{
			Input:  5,
			Output: 5,
		},
		{
			Input:  10,
			Output: 55,
		},
		{
			Input:  15,
			Output: 610,
		},
		{
			Input:  20,
			Output: 6765,
		},
		{
			Input:  25,
			Output: 75025,
		},
		{
			Input:  25,
			Output: 75025,
		},
		{
			Input:  30,
			Output: 832040,
		},
		{
			Input:  35,
			Output: 9227465,
		},
		{
			Input:  40,
			Output: 102334155,
		},
	}

	for _, test := range tests {
		out := FibonacciBottomUp(test.Input)
		assert.Assert(t, out == test.Output, "input: %d, expected: %d, got: %d", test.Input, test.Output, out)
	}
}
