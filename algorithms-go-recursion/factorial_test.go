package recursion

import (
	"testing"

	"gotest.tools/assert"
)

func TestFactorial(t *testing.T) {
	var got int64

	inputOutput := map[int64]int64{
		0:  1,
		1:  1,
		2:  2,
		3:  6,
		4:  24,
		5:  120,
		6:  720,
		7:  5040,
		8:  40320,
		9:  362880,
		10: 3628800,
		11: 39916800,
		12: 479001600,
		13: 6227020800,
		14: 87178291200,
		15: 1307674368000,
		16: 20922789888000,
		17: 355687428096000,
		18: 6402373705728000,
		19: 121645100408832000,
		20: 2432902008176640000,
		21: -4249290049419214848,
	}

	for input, output := range inputOutput {
		got = IterativeFactorial(input)
		assert.Assert(t, got == output, "got: %d", got)
		got = Factorial(input)
		assert.Assert(t, got == output, "got: %d", got)
	}
}
