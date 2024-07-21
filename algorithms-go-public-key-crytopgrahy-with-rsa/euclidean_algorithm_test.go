package cryptography

import (
	"testing"

	"gotest.tools/assert"
)

func TestGDC(t *testing.T) {
	got := GDC(270, 192)
	assert.Assert(t, got == 6)

	got = GDC(7469, 2464)
	assert.Assert(t, got == 77)

	got = GDC(55290, 115430)
	assert.Assert(t, got == 970)
}

func TestLCM(t *testing.T) {
	got := LCM(270, 192)
	assert.Assert(t, got == 8640)

	got = LCM(7469, 2464)
	assert.Assert(t, got == 239008)

	got = LCM(55290, 115430)
	assert.Assert(t, got == 6579510)
}
