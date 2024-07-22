package cryptography

import (
	"testing"

	"gotest.tools/assert"
)

func TestFindFactors(t *testing.T) {
	expected := []int{2, 5}
	got := FindFactors(10)
	assert.DeepEqual(t, expected, got)

	got = FindFactors(312680865509917)
	expected = []int{7791799, 40129483}
	assert.DeepEqual(t, expected, got)

	got = FindFactors(12345678901234)
	expected = []int{2, 7, 73, 12079920647}
	assert.DeepEqual(t, expected, got)

	got = FindFactors(64374108854777)
	expected = []int{64374108854777}
	assert.DeepEqual(t, expected, got)
}

func TestFindFactorsSieve(t *testing.T) {
	primes = SieveToPrimes(EulerSieve(20000000))
	expected := []int{2, 5}
	got := FindFactorsSieve(10)
	assert.DeepEqual(t, expected, got)

	got = FindFactorsSieve(312680865509917)
	expected = []int{7791799, 40129483}
	assert.DeepEqual(t, expected, got)

	got = FindFactorsSieve(12345678901234)
	expected = []int{2, 7, 73, 12079920647}
	assert.DeepEqual(t, expected, got)

	got = FindFactorsSieve(64374108854777)
	expected = []int{64374108854777}
	assert.DeepEqual(t, expected, got)
}
