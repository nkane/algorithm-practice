package hash_table

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestLinearProbingHashTable(t *testing.T) {
	employees := []Employee{
		{"Ann Archer", "202-555-0101"},
		{"Bob Baker", "202-555-0102"},
		{"Cindy Cant", "202-555-0103"},
		{"Dan Deever", "202-555-0104"},
		{"Edwina Eager", "202-555-0105"},
		{"Fred Franklin", "202-555-0106"},
		{"Gina Gable", "202-555-0107"},
	}
	ht := NewLinearProbingHashTable(10)
	for _, e := range employees {
		ht.Set(e.Name, e.Phone)
	}
	ht.Dump()
	assert.Assert(t, ht.Contains("Sally Owens") == false)
	assert.Assert(t, ht.Contains("Dan Deever") == true)
	assert.Assert(t, ht.Get("Fred Franklin") == "202-555-0106")
	ht.Set("Fred Franklin", "202-555-0100")
	assert.Assert(t, ht.Get("Fred Franklin") == "202-555-0100")

	fmt.Println(time.Now())                   // Print the time so it will compile if we use a fixed seed.
	random := rand.New(rand.NewSource(12345)) // Initialize with a fixed seed
	bigCapacity := 1009
	bigHashTable := NewLinearProbingHashTable(bigCapacity)
	numItems := int(float32(bigCapacity) * 0.9)
	for i := 0; i < numItems; i++ {
		str := fmt.Sprintf("%d-%d", i, random.Intn(1000000))
		bigHashTable.Set(str, str)
	}
	bigHashTable.DumpConise()
	fmt.Printf("Average probe sequence length: %f\n", bigHashTable.AveProbeSequenceLength())
}
