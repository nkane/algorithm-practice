package cryptography

import (
	"math/rand"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func Totient(p int, q int) int {
	// return (p-1) * (q -1) // Euler's totient function
	return LCM(p-1, q-1) // Carmicheal's totient function
}

func RandomExponent(n int) int {
	for {
		e := RandRange(3, n)
		if GDC(e, n) == 1 {
			return e
		}
	}
}

func InverseMod(a int, modulus int) int {
	t := 0
	newt := 1
	r := modulus
	newr := a
	for newr != 0 {
		quotient := r / newr
		t, newt = newt, t-quotient*newt
		r, newr = newr, r-quotient*newr
	}
	if r > 1 {
		return -1
	}
	if t < 0 {
		t = t + modulus
	}
	return t
}
