package hash_table

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestDoubleHashTable(t *testing.T) {
	employees := []Employee{
		{"Ann Archer", "202-555-0101", false},
		{"Bob Baker", "202-555-0102", false},
		{"Cindy Cant", "202-555-0103", false},
		{"Dan Deever", "202-555-0104", false},
		{"Edwina Eager", "202-555-0105", false},
		{"Fred Franklin", "202-555-0106", false},
		{"Gina Gable", "202-555-0107", false},
	}
	ht := NewDoubleHashTable(10)
	for _, e := range employees {
		ht.Set(e.Name, e.Phone)
	}

	ht.Dump()
	assert.Assert(t, ht.Contains("Sally Owens") == false)
	assert.Assert(t, ht.Contains("Dan Deever") == true)
	ht.Delete("Dan Deever")
	assert.Assert(t, ht.Contains("Dan Deever") == false)

	assert.Assert(t, ht.Get("Fred Franklin") == "202-555-0106")
	ht.Set("Fred Franklin", "202-555-0100")
	assert.Assert(t, ht.Get("Fred Franklin") == "202-555-0100")
	ht.Dump()
}

func TestDoubleHashTable_Probe(t *testing.T) {
	employees := []Employee{
		{"Ann Archer", "202-555-0101", false},
		{"Bob Baker", "202-555-0102", false},
		{"Cindy Cant", "202-555-0103", false},
		{"Dan Deever", "202-555-0104", false},
		{"Edwina Eager", "202-555-0105", false},
		{"Fred Franklin", "202-555-0106", false},
		{"Gina Gable", "202-555-0107", false},
	}
	ht := NewDoubleHashTable(10)
	for _, e := range employees {
		ht.Set(e.Name, e.Phone)
	}
	ht.Dump()
	ht.Probe("Hank Hardy")
	ht.Probe("Ann Archer")
	ht.Probe("Cindy Cant")
	ht.Probe("Dan Deever")
	ht.Probe("Edwina Eagar")
	ht.Probe("Fred Franklin")
	ht.Probe("Gina Gable")
	ht.Set("Hank Hardy", "202-555-0108")
	ht.Probe("Hank Hardy")
}

func TestDoubleHashTable_Big(t *testing.T) {
	fmt.Println(time.Now())                   // Print the time so it will compile if we use a fixed seed.
	random := rand.New(rand.NewSource(12345)) // Initialize with a fixed seed
	bigCapacity := 1009
	bigHashTable := NewDoubleHashTable(bigCapacity)
	numItems := int(float32(bigCapacity) * 0.9)
	for i := 0; i < numItems; i++ {
		str := fmt.Sprintf("%d-%d", i, random.Intn(1000000))
		bigHashTable.Set(str, str)
	}
	bigHashTable.DumpConise()
	fmt.Printf("Average probe sequence length: %f\n", bigHashTable.AveProbeSequenceLength())
}
