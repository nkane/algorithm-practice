package hash_table

import (
	"testing"

	"gotest.tools/assert"
)

func TestChainingHashTable(t *testing.T) {
	employees := []Employee{
		{"Ann Archer", "202-555-0101"},
		{"Bob Baker", "202-555-0102"},
		{"Cindy Cant", "202-555-0103"},
		{"Dan Deever", "202-555-0104"},
		{"Edwina Eager", "202-555-0105"},
		{"Fred Franklin", "202-555-0106"},
		{"Gina Gable", "202-555-0107"},
		{"Herb Henshaw", "202-555-0108"},
		{"Ida Iverson", "202-555-0109"},
		{"Jeb Jacobs", "202-555-0110"},
	}

	ht := NewChainingHashTable(10)
	for _, employee := range employees {
		ht.Set(employee.Name, employee.Phone)
	}
	ht.Dump()
	assert.Assert(t, ht.Contains("Sally Ownens") == false)
	assert.Assert(t, ht.Contains("Dan Deever") == true)
	ht.Delete("Dan Deever")
	assert.Assert(t, ht.Contains("Fred Franklin") == true)
	ht.Set("Fred Franklin", "202-555-0100")
	assert.Assert(t, ht.Get("Fred Franklin") == "202-555-0100")
}
