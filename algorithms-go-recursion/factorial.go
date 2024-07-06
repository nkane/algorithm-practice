package recursion

func IterativeFactorial(n int64) int64 {
	result := int64(1)
	for i := int64(2); i <= n; i++ {
		result *= i
	}
	return result
}

func Factorial(n int64) int64 {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}
