package cryptography

import (
	"fmt"
	"math"
	"testing"
)

func TestPrimality(t *testing.T) {
	primes := []int{
		10009, 11113, 11699, 12809, 14149,
		15643, 17107, 17881, 19301, 19793,
	}
	composites := []int{
		10323, 11397, 12212, 13503, 14599,
		16113, 17547, 17549, 18893, 19999,
	}

	fmt.Printf("Probability: %f%%\n\n", (1.0-1.0/math.Pow(2, numTests))*100.0)

	fmt.Println("Primes:")
	for _, number := range primes {
		if IsProbablyPrime(number, 10) {
			fmt.Println(number, " Prime")
		} else {
			fmt.Println(number, " Composite")
		}
	}
	fmt.Println()

	fmt.Println("Composites:")
	for _, number := range composites {
		if IsProbablyPrime(number, 10) {
			fmt.Println(number, " Prime")
		} else {
			fmt.Println(number, " Composite")
		}
	}
}
