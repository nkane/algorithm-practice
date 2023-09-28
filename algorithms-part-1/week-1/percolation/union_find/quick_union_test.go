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

// TODO(nick): write a test for 4 elements
