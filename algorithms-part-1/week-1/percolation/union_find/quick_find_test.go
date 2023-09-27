package union_find

import (
	"math/rand"
	"testing"
)

func TestQuickFindCreate(t *testing.T) {
	n := rand.Intn(1024)
	uf := CreateQuickFind(n)
	ids := uf.IDs()
	for i, v := range ids {
		if v != i {
			t.Fatalf("failed at id comparison, i: %d, v: %d\n", i, v)
		}
	}
	if uf.Count() != n {
		t.Fatalf("expected count to be equal to n: %d, got: %d\n", n, uf.Count())
	}
}

func TestQuickFindTwoElements(t *testing.T) {
	expectedIDs := []int{1, 1}
	uf := CreateQuickFind(2)
	uf.Union(0, 1)
	if uf.Count() != 1 {
		t.Fatalf("expected count to be 1, got: %d\n", uf.Count())
	}
	if !CompareIDs(expectedIDs, uf.IDs()) {
		t.Fatalf("expected ids to be same\nexpected: %+v\ngot:%+v\n", expectedIDs, uf.IDs())
	}
	if !uf.Connected(0, 1) {
		t.Fatalf("expected 0 and 1 to be connected, is not connected")
	}
}

func TestQuickFindTinyInput(t *testing.T) {
	expectedOutput := []int{1, 1, 1, 8, 8, 1, 1, 1, 8, 8}
	uf, err := CreateQuickFindFromFile("./data/tinyUF.txt")
	if err != nil {
		t.Fatalf("failed to create quick find\n")
	}
	if uf.Count() != 2 {
		t.Fatalf("expected component size of 2, got: %d\n", uf.Count())
	}
	if !CompareIDs(expectedOutput, uf.IDs()) {
		t.Fatalf("expected ids to be the same\n")
	}
}

func CompareIDs(expectedIDs []int, ids []int) bool {
	if len(expectedIDs) != len(ids) {
		return false
	}
	for idx, v := range ids {
		if expectedIDs[idx] != v {
			return false
		}
	}
	return true
}
