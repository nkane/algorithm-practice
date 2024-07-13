package hash_table

import (
	"fmt"
)

type Employee struct {
	Name  string
	Phone string
}

type ChainingHashTable struct {
	NumBuckets int
	Buckets    [][]*Employee
}

func NewChainingHashTable(numBuckets int) *ChainingHashTable {
	return &ChainingHashTable{
		NumBuckets: numBuckets,
		Buckets:    make([][]*Employee, numBuckets),
	}
}

func (ht *ChainingHashTable) Dump() {
	for i, bucket := range ht.Buckets {
		fmt.Printf("Bucket %d\n", i)
		for _, item := range bucket {
			fmt.Printf("\t%s: %s\n", item.Name, item.Phone)
		}
	}
}

func (ht *ChainingHashTable) Hash(value string) int {
	hash := 5381
	for _, ch := range value {
		hash = ((hash << 5) + hash) + int(ch)
	}
	if hash < 0 {
		hash = -hash
	}
	return hash % ht.NumBuckets
}

func (ht *ChainingHashTable) Find(name string) (int, int) {
	keyIdx := ht.Hash(name)
	bucket := ht.Buckets[keyIdx]
	dataIdx := -1
	for i, item := range bucket {
		if item.Name == name {
			dataIdx = i
			break
		}
	}
	return keyIdx, dataIdx
}

func (ht *ChainingHashTable) Set(name string, phone string) {
	keyIdx, dataIdx := ht.Find(name)
	if dataIdx >= 0 {
		ht.Buckets[keyIdx][dataIdx].Phone = phone
	} else {
		ht.Buckets[keyIdx] = append(ht.Buckets[keyIdx], &Employee{
			Name:  name,
			Phone: phone,
		})
	}
}

func (ht *ChainingHashTable) Get(name string) string {
	keyIdx, dataIdx := ht.Find(name)
	if dataIdx >= 0 {
		return ht.Buckets[keyIdx][dataIdx].Phone
	}
	return ""
}

func (ht *ChainingHashTable) Contains(name string) bool {
	_, dataIdx := ht.Find(name)
	return dataIdx >= 0
}

func (ht *ChainingHashTable) Delete(name string) error {
	keyIdx, dataIdx := ht.Find(name)
	if dataIdx >= 0 {
		ht.Buckets[keyIdx] = append(ht.Buckets[keyIdx][:dataIdx], ht.Buckets[keyIdx][dataIdx+1:]...)
	} else {
		return fmt.Errorf("hash table doesn't not contain key: %s", name)
	}
	return nil
}
