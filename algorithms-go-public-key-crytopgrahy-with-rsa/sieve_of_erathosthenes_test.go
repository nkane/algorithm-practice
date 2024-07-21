package cryptography

import (
	"bytes"
	"os"
	"testing"

	"gotest.tools/assert"
)

func TestSieveOfEratosthenes(t *testing.T) {
	expected := []bool{
		false, false, true, true, false, true, false, true,
		false, false, false, true, false, true, false, false,
		false,
	}
	got := SieveOfEratosthenes(16)
	assert.Assert(t, len(expected) == len(got))
	assert.DeepEqual(t, expected, got)
}

func TestSieveOfEratosthenesPrint(t *testing.T) {
	got := SieveOfEratosthenes(16)

	oldOut := os.Stdout
	r, w, err := os.Pipe()
	assert.NilError(t, err)
	os.Stdout = w

	PrintSieve(got)

	w.Close()
	os.Stdout = oldOut
	buf := bytes.Buffer{}
	buf.ReadFrom(r)
	out := buf.String()
	expected := "2 3 5 7 11 13\n"
	assert.Assert(t, out == expected, "\nexpected: %s\ngot: %s", expected, got)
}

func TestSieveOfEratosthenesToPrimes(t *testing.T) {
	expected := []int{2, 3, 5, 7, 11, 13}
	sieve := SieveOfEratosthenes(16)
	got := SieveToPrimes(sieve)
	assert.DeepEqual(t, expected, got)
}
