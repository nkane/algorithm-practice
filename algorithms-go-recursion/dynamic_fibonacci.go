package recursion

// NOTE(nick): the able has to start with some values to avoid nil deref, this wasn't
// described in the project
var fibonacciValues = []int64{0, 1}

// approach #1
func FibonacciOnTheFly(n int64) int64 {
	if n >= int64(len(fibonacciValues)) {
		fibonacciValues = append(fibonacciValues, FibonacciOnTheFly(n-1)+FibonacciOnTheFly(n-2))
	}
	return fibonacciValues[n]
}

// approach #2
func InitializeSlice() {
	fibonacciValues := make([]int64, 93)
	fibonacciValues[0] = 0
	fibonacciValues[1] = 1
	for i := 2; i < len(fibonacciValues); i++ {
		fibonacciValues[i] = fibonacciValues[i-1] + fibonacciValues[i-2]
	}
}

func FibonacciPrefilled(n int64) int64 {
	// NOTE(nick): requires the slice to be prefilled
	return fibonacciValues[n]
}

// approach #3 without prefilled slice
func FibonacciBottomUp(n int64) int64 {
	if n <= 1 {
		return n
	}
	fibIMinus2 := int64(0)
	fibIMinus1 := int64(1)
	fibI := fibIMinus1 + fibIMinus2
	for i := int64(1); i < n; i++ {
		// calculate the value of fibI
		fibI = fibIMinus1 + fibIMinus2
		// set fibIMinus2 and fibIMinus1 for the next value
		fibIMinus2 = fibIMinus1
		fibIMinus1 = fibI
	}
	return fibI
}
