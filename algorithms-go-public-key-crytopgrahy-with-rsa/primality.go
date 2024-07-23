package cryptography

import (
	"fmt"
)

const numTests = 20

// var random = rand.New(rand.NewSource(12346))
// var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandRange(min int, max int) int {
	return min + random.Intn(max-min)
}

func IsProbablyPrime(p int, numTests int) bool {
	for i := 0; i < numTests; i++ {
		n := RandRange(1, p)
		result := FastExpMod(n, p-1, p)
		if result < 0 {
			fmt.Println("interger overflow")
		}
		if result != 1 {
			return false
		}
	}
	return true
}

func FindPrime(min int, max int, numTests int) int {
	for {
		p := RandRange(min, max)
		if p%2 == 0 {
			continue
		}
		if IsProbablyPrime(p, numTests) {
			return p
		}
	}
}
