package algorithms

import (
	"math/rand"
	"testing"

	"gotest.tools/assert"
)

func TestLinearSearch(t *testing.T) {
	length := 100
	slice := MakeRandomSlice(length, 50)
	got, _ := LinearSearch(slice, slice[rand.Intn(length)])
	assert.Assert(t, got > -1)
}
