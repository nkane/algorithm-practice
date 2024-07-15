package hash_table

import (
	"fmt"
)

type QuadraticProbingHashTable struct {
	Capacity  int
	Employees []*Employee
}

func NewQuadraticProbingHashTable(capacity int) *QuadraticProbingHashTable {
	return &QuadraticProbingHashTable{
		Capacity:  capacity,
		Employees: make([]*Employee, capacity),
	}
}

func (ht *QuadraticProbingHashTable) Dump() {
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

func (ht *QuadraticProbingHashTable) Hash(value string) int {
	hash := 5381
	for _, ch := range value {
		hash = ((hash << 5) + hash) + int(ch)
	}
	if hash < 0 {
		hash = -hash
	}
	return hash
}

func (ht *QuadraticProbingHashTable) Find(name string) (int, int) {
	hash := ht.Hash(name) % ht.Capacity
	deletedIndex := -1
	for i := 0; i < ht.Capacity; i++ {
		index := (hash + (i * i)) % ht.Capacity
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

func (ht *QuadraticProbingHashTable) Set(name string, phone string) error {
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

func (ht *QuadraticProbingHashTable) Get(name string) string {
	index, _ := ht.Find(name)
	if index < 0 ||
		ht.Employees[index] == nil ||
		ht.Employees[index].Deleted {
		return ""
	}
	return ht.Employees[index].Phone
}

func (ht *QuadraticProbingHashTable) Contains(name string) bool {
	index, _ := ht.Find(name)
	if index < 0 ||
		ht.Employees[index] == nil ||
		ht.Employees[index].Deleted {
		return false
	}
	return ht.Employees[index].Name == name
}

func (ht *QuadraticProbingHashTable) Delete(name string) {
	index, _ := ht.Find(name)
	if index >= 0 &&
		ht.Employees[index] != nil &&
		!ht.Employees[index].Deleted {
		ht.Employees[index].Deleted = true
	}
}

func (ht *QuadraticProbingHashTable) DumpConise() {
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

func (ht *QuadraticProbingHashTable) AveProbeSequenceLength() float32 {
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

func (ht *QuadraticProbingHashTable) Probe(name string) int {
	hash := ht.Hash(name) % ht.Capacity
	fmt.Printf("Probing %s (%d)\n", name, hash)
	deletedItem := -1
	for i := 0; i < ht.Capacity; i++ {
		index := (hash + (i * i)) % ht.Capacity
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
