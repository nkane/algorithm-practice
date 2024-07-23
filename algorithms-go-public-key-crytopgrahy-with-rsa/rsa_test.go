package cryptography

import (
	"fmt"
	"testing"
)

func TestRSA(t *testing.T) {
	const numtTests = 100
	// pick two prime numbers
	p := FindPrime(10000, 50000, numtTests)
	q := FindPrime(10000, 50000, numtTests)

	// calculate the public key modulus n
	n := p * q

	// calculate carmicheal's totient function
	lambdaN := Totient(p, q)

	// pick a random public key exponent e in the range [3, lambda_n]
	// where gdc(e, lambda_n) = 1
	e := RandomExponent(lambdaN)

	// find the inverse of e mod lambda_n
	d := InverseMod(e, lambdaN)

	fmt.Printf("*** Public ***\n")
	fmt.Printf("Public key modulus:    %d\n", n)
	fmt.Printf("Public key exponent e: %d\n", e)
	fmt.Printf("\n*** Private ***\n")
	fmt.Printf("Primes:    %d, %d\n", p, q)
	fmt.Printf("λ(n):      %d\n", lambdaN)
	fmt.Printf("d:         %d\n", d)

	m := 128
	ciphertext := FastExpMod(m, e, n)
	fmt.Printf("Ciphertext: %d\n", ciphertext)

	plaintext := FastExpMod(ciphertext, d, n)
	fmt.Printf("Plaintext:  %d\n", plaintext)
}
