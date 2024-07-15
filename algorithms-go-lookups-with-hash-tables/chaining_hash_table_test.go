package hash_table

import (
	"testing"

	"gotest.tools/assert"
)

func TestChainingHashTable(t *testing.T) {
	employees := []Employee{
		{"Ann Archer", "202-555-0101", false},
		{"Bob Baker", "202-555-0102", false},
		{"Cindy Cant", "202-555-0103", false},
		{"Dan Deever", "202-555-0104", false},
		{"Edwina Eager", "202-555-0105", false},
		{"Fred Franklin", "202-555-0106", false},
		{"Gina Gable", "202-555-0107", false},
		{"Herb Henshaw", "202-555-0108", false},
		{"Ida Iverson", "202-555-0109", false},
		{"Jeb Jacobs", "202-555-0110", false},
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
