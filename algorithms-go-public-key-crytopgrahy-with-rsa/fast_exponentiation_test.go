package cryptography

import (
	"testing"

	"gotest.tools/assert"
)

func TestFastExp(t *testing.T) {
	got := FastExp(8, 6)
	assert.Assert(t, got == 262144)

	got = FastExp(7, 10)
	assert.Assert(t, got == 282475249)

	got = FastExp(9, 13)
	assert.Assert(t, got == 2541865828329)

	got = FastExp(213, 5)
	assert.Assert(t, got == 438427732293)
}

func TestFastExpMod(t *testing.T) {
	got := FastExpMod(8, 6, 10)
	assert.Assert(t, got == 4)

	got = FastExpMod(7, 10, 101)
	assert.Assert(t, got == 65)

	got = FastExpMod(9, 13, 283)
	assert.Assert(t, got == 179)

	got = FastExpMod(213, 5, 1000)
	assert.Assert(t, got == 293)
}
