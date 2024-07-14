package hash_table

import "fmt"

type LinearProbingHashTable struct {
	Capacity  int
	Employees []*Employee
}

func NewLinearProbingHashTable(capacity int) *LinearProbingHashTable {
	return &LinearProbingHashTable{
		Capacity:  capacity,
		Employees: make([]*Employee, capacity),
	}
}

func (ht *LinearProbingHashTable) Dump() {
	for i, e := range ht.Employees {
		if e != nil {
			fmt.Printf("%d: %s\t%s\n", i, e.Name, e.Phone)
		} else {
			fmt.Printf("%d: ---\n", i)
		}
	}
	fmt.Println()
}

func (ht *LinearProbingHashTable) Hash(value string) int {
	hash := 5381
	for _, ch := range value {
		hash = ((hash << 5) + hash) + int(ch)
	}
	if hash < 0 {
		hash = -hash
	}
	return hash
}

func (ht *LinearProbingHashTable) Find(name string) (int, int) {
	hash := ht.Hash(name) % ht.Capacity
	for i := 0; i < ht.Capacity; i++ {
		index := (hash + i) % ht.Capacity
		if ht.Employees[index] == nil ||
			ht.Employees[index].Name == name {
			return index, i + 1
		}
	}
	return -1, ht.Capacity
}

func (ht *LinearProbingHashTable) Set(name string, phone string) error {
	index, _ := ht.Find(name)
	if index < 0 {
		return fmt.Errorf("key is not in table")
	}
	if ht.Employees[index] == nil {
		ht.Employees[index] = &Employee{
			Name:  name,
			Phone: phone,
		}
	} else {
		ht.Employees[index].Name = name
		ht.Employees[index].Phone = phone
	}
	return nil
}

func (ht *LinearProbingHashTable) Get(name string) string {
	index, _ := ht.Find(name)
	if index < 0 || ht.Employees[index] == nil {
		return ""
	}
	return ht.Employees[index].Phone
}

func (ht *LinearProbingHashTable) Contains(name string) bool {
	index, _ := ht.Find(name)
	if index < 0 || ht.Employees[index] == nil {
		return false
	}
	return true
}

func (ht *LinearProbingHashTable) DumpConise() {
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

func (ht *LinearProbingHashTable) AveProbeSequenceLength() float32 {
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
