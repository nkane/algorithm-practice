package union_find

import (
	"math/rand"
	"testing"
)

func TestQuickUnionCreate(t *testing.T) {
	n := rand.Intn(1024)
	uf := CreateQuickUnion(n)
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

func TestQuickUnionTwoElements(t *testing.T) {
	expectedIDs := []int{1, 1}
	uf := CreateQuickUnion(2)
	uf.Union(0, 1)
	if uf.Count() != 1 {
		t.Fatalf("expected count to be 1, got: %d\n", uf.Count())
	}
	if !uf.Connected(0, 1) {
		t.Fatalf("expected 0 and 1 to be connected, is not connected")
	}
	if !CompareIDs(expectedIDs, uf.IDs()) {
		t.Fatalf("expected ids to be same\nexpected: %+v\ngot:%+v\n", expectedIDs, uf.IDs())
	}
}

func TestQuickUnionFourElements(t *testing.T) {
	expectedIDs := []int{3, 1, 2, 3}
	uf := CreateQuickUnion(4)
	uf.Union(0, 3)
	if uf.Count() != 3 {
		t.Fatalf("expected count to be 3, got: %d\n", uf.Count())
	}
	if !CompareIDs(expectedIDs, uf.IDs()) {
		t.Fatalf("expected ids to be same\nexpected: %+v\ngot:%+v\n", expectedIDs, uf.IDs())
	}
	uf.Union(3, 1)
	if uf.Count() != 2 {
		t.Fatalf("expected count to be 2, got: %d\n", uf.Count())
	}
	expectedIDs = []int{3, 1, 2, 1}
	if !CompareIDs(expectedIDs, uf.IDs()) {
		t.Fatalf("expected ids to be same\nexpected: %+v\ngot:%+v\n", expectedIDs, uf.IDs())
	}
	uf.Union(2, 3)
	if uf.Count() != 1 {
		t.Fatalf("expected count to be 1, got: %d\n", uf.Count())
	}
	expectedIDs = []int{3, 1, 1, 1}
	if !CompareIDs(expectedIDs, uf.IDs()) {
		t.Fatalf("expected ids to be same\nexpected: %+v\ngot:%+v\n", expectedIDs, uf.IDs())
	}
}

func TestQuickUnionTinyInput(t *testing.T) {
	expectedIDs := []int{1, 1, 1, 8, 3, 0, 5, 1, 8, 8}
	uf, err := CreateQuickUnionFromFile("./data/tinyUF.txt")
	if err != nil {
		t.Fatalf("failed to create quick find\n")
	}
	if uf.Count() != 2 {
		t.Fatalf("expected component size of 2, got: %d\n", uf.Count())
	}
	if !CompareIDs(expectedIDs, uf.IDs()) {
		t.Fatalf("expected ids to be same\nexpected: %+v\ngot:%+v\n", expectedIDs, uf.IDs())
	}
}
