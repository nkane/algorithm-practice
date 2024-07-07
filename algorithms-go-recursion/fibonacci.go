package recursion

func Fibonacci(n int64) int64 {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
