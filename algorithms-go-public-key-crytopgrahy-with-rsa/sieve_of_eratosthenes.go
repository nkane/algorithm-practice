package cryptography

import (
	"fmt"
)

func SieveOfEratosthenes(max int) []bool {
	output := make([]bool, max+1)
	output[2] = true
	for i := 3; i < len(output); i += 2 {
		output[i] = true
	}
	for i := 3; i <= max; i += 2 {
		if output[i] {
			for j := i * 3; j <= max; j += i {
				output[j] = false
			}
		}
	}
	return output
}

func PrintSieve(sieve []bool) {
	fmt.Print("2")
	for i := 3; i < len(sieve); i += 2 {
		if sieve[i] {
			fmt.Print(" ", i)
		}
	}
	fmt.Println()
}

func SieveToPrimes(sieve []bool) []int {
	primes := []int{2}
	for i := 3; i < len(sieve); i += 2 {
		if sieve[i] {
			primes = append(primes, i)
		}
	}
	return primes
}
