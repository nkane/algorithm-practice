package hash_table

import (
	"fmt"
)

type DoubleHashTable struct {
	Capacity  int
	Employees []*Employee
}

func NewDoubleHashTable(capacity int) *DoubleHashTable {
	return &DoubleHashTable{
		Capacity:  capacity,
		Employees: make([]*Employee, capacity),
	}
}

func (ht *DoubleHashTable) Dump() {
	for i, e := range ht.Employees {
		if e != nil {
			if e.Deleted {
				fmt.Printf("%d: xxx\n", i)
			} else {
				fmt.Printf("%d: %s\t%s\n", i, e.Name, e.Phone)
			}
		} else {
			fmt.Printf("%d: ---\n", i)
		}
	}
	fmt.Println()
}

func (ht *DoubleHashTable) Hash(value string) int {
	hash := 5381
	for _, ch := range value {
		hash = ((hash << 5) + hash) + int(ch)
	}
	if hash < 0 {
		hash = -hash
	}
	return hash
}

func (ht *DoubleHashTable) Hash2(value string) int {
	hash := 0
	for _, ch := range value {
		hash += int(ch)
		hash += hash << 10
		hash ^= hash >> 6
	}
	if hash < 0 {
		hash = -hash
	}
	if hash == 0 {
		hash = 1
	}
	return hash
}

func (ht *DoubleHashTable) Find(name string) (int, int) {
	hash1 := ht.Hash(name) % ht.Capacity
	hash2 := ht.Hash2(name) % ht.Capacity
	deletedIndex := -1
	for i := 0; i < ht.Capacity; i++ {
		index := (hash1 + i*hash2) % ht.Capacity
		if ht.Employees[index] == nil {
			if deletedIndex >= 0 {
				return deletedIndex, i + 1
			}
			return index, i + 1
		}
		if ht.Employees[index].Deleted {
			if deletedIndex < 0 {
				deletedIndex = index
			}
		} else if ht.Employees[index].Name == name {
			return index, i + 1
		}
	}
	if deletedIndex >= 0 {
		return deletedIndex, ht.Capacity
	}
	return -1, ht.Capacity
}

func (ht *DoubleHashTable) Set(name string, phone string) error {
	index, _ := ht.Find(name)
	if index < 0 {
		return fmt.Errorf("key is not in table")
	}
	if ht.Employees[index] == nil || ht.Employees[index].Deleted {
		ht.Employees[index] = &Employee{
			Name:  name,
			Phone: phone,
		}
	} else {
		ht.Employees[index].Phone = phone
	}
	return nil
}

func (ht *DoubleHashTable) Get(name string) string {
	index, _ := ht.Find(name)
	if index < 0 ||
		ht.Employees[index] == nil ||
		ht.Employees[index].Deleted {
		return ""
	}
	return ht.Employees[index].Phone
}

func (ht *DoubleHashTable) Contains(name string) bool {
	index, _ := ht.Find(name)
	if index < 0 ||
		ht.Employees[index] == nil ||
		ht.Employees[index].Deleted {
		return false
	}
	return ht.Employees[index].Name == name
}

func (ht *DoubleHashTable) Delete(name string) {
	index, _ := ht.Find(name)
	if index >= 0 &&
		ht.Employees[index] != nil &&
		!ht.Employees[index].Deleted {
		ht.Employees[index].Deleted = true
	}
}

func (ht *DoubleHashTable) DumpConise() {
	for i, e := range ht.Employees {
		if e == nil {
			fmt.Printf(".")
		} else {
			fmt.Printf("O")
		}
		if i%50 == 49 {
			fmt.Println()
		}
	}
	fmt.Println()
}

func (ht *DoubleHashTable) AveProbeSequenceLength() float32 {
	totalLength := 0
	numValues := 0
	for _, e := range ht.Employees {
		if e != nil {
			_, probeLength := ht.Find(e.Name)
			totalLength += probeLength
			numValues++
		}
	}
	return float32(totalLength) / float32(numValues)
}

func (ht *DoubleHashTable) Probe(name string) int {
	hash1 := ht.Hash(name) % ht.Capacity
	hash2 := ht.Hash2(name) % ht.Capacity
	fmt.Printf("Probing %s (%d, %d)\n", name, hash1, hash2)
	deletedItem := -1
	for i := 0; i < ht.Capacity; i++ {
		index := (hash1 + i*hash2) % ht.Capacity
		fmt.Printf("\t%d: ", index)
		if ht.Employees[index] == nil {
			fmt.Printf("---\n")
		} else if ht.Employees[index].Deleted {
			fmt.Printf("xxx\n")
		} else {
			fmt.Printf("%s\n", ht.Employees[index].Name)
		}
		if ht.Employees[index] == nil {
			if deletedItem >= 0 {
				fmt.Printf("\tReturning deleted index %d\n", deletedItem)
				return deletedItem
			}
			fmt.Printf("\tReturning nil index %d\n", index)
			return index
		}
		if ht.Employees[index].Deleted {
			if deletedItem < 0 {
				deletedItem = index
			}
		} else if ht.Employees[index].Name == name {
			fmt.Printf("\tReturn found index %d\n", index)
			return index
		}
	}
	if deletedItem >= 0 {
		fmt.Printf("\tReturning deleted index %d\n", deletedItem)
		return deletedItem
	}
	fmt.Printf("\tTable is full\n")
	return -1
}
