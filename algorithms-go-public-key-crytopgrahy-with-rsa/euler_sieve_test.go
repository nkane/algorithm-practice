package cryptography

import (
	"bytes"
	"os"
	"testing"

	"gotest.tools/assert"
)

func TestEulerSieve(t *testing.T) {
	expected := []bool{
		false, false, true, true, false, true, false, true,
		false, false, false, true, false, true, false, false,
		false,
	}
	got := EulerSieve(16)
	assert.Assert(t, len(expected) == len(got))
	assert.DeepEqual(t, expected, got)
}

func TestEulerSievePrint(t *testing.T) {
	oldOut := os.Stdout
	r, w, err := os.Pipe()
	assert.NilError(t, err)
	os.Stdout = w

	got := EulerSieve(16)
	PrintSieve(got)

	w.Close()
	os.Stdout = oldOut
	buf := bytes.Buffer{}
	buf.ReadFrom(r)
	out := buf.String()
	expected := "2 3 5 7 11 13\n"
	assert.Assert(t, out == expected, "\nexpected: %s\ngot: %s", expected, got)
}

func TestEulerSieveToPrimes(t *testing.T) {
	expected := []int{2, 3, 5, 7, 11, 13}
	sieve := EulerSieve(16)
	got := SieveToPrimes(sieve)
	assert.DeepEqual(t, expected, got)
}
