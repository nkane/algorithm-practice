package binary_search

import (
	"testing"

	"gotest.tools/v3/assert"
)

type Test struct {
	List               []int
	Find               int
	ExpectedFound      bool
	ExpectedIterations int
}

func TestBinarySearch(t *testing.T) {
	tests := []Test{
		{
			List: []int{
				1,
				2,
				3,
				4,
				5,
				6,
			},
			Find:               3,
			ExpectedFound:      true,
			ExpectedIterations: 1,
		},
		{
			List: []int{
				1,
				2,
				3,
				4,
				5,
				6,
			},
			Find:               5,
			ExpectedFound:      true,
			ExpectedIterations: 2,
		},
		{
			List: []int{
				1,
				2,
				3,
				4,
				5,
				6,
			},
			Find:               4,
			ExpectedFound:      true,
			ExpectedIterations: 3,
		},
		{
			List: []int{
				1,
				2,
				3,
				4,
				5,
				6,
			},
			Find:               1,
			ExpectedFound:      true,
			ExpectedIterations: 2,
		},
		{
			List: []int{
				1,
				2,
				3,
				4,
				5,
				6,
			},
			Find:               2,
			ExpectedFound:      true,
			ExpectedIterations: 3,
		},
		{
			List: []int{
				1,
				2,
				3,
				4,
				5,
				6,
				7,
				8,
				9,
				10,
			},
			Find:               2,
			ExpectedFound:      true,
			ExpectedIterations: 3,
		},
	}
	for _, test := range tests {
		sr := BinarySearch(test.List, test.Find)
		assert.Assert(t, test.ExpectedFound == sr.Found, "expected: %+v, got: %+v", test.ExpectedFound, sr.Found)
		assert.Assert(t, test.ExpectedIterations == sr.Iterations, "expected: %+v, got: %+v", test.ExpectedIterations, sr.Iterations)
	}
}
