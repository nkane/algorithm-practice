package union_find

import (
	"math/rand"
	"testing"
)

func TestQuickWeightedUnionCreate(t *testing.T) {
	n := rand.Intn(1024)
	uf := CreateQuickWeightedUnion(n)
	// cast to type to check fields
	qwu, ok := (uf).(*QuickWeightedUnion)
	if !ok {
		t.Fatal("created type is not expected type of QuickWeightedUnion")
	}
	if len(qwu.Elements) != len(qwu.Sizes) {
		t.Fatal("expected elements and size slices to be the same length")
	}
	for i := 0; i < len(qwu.Elements); i++ {
		if qwu.Elements[i] != i {
			t.Fatalf("expected element to be equal to %d, got %d", i, qwu.Elements[i])
		}
		if qwu.Sizes[i] != 1 {
			t.Fatalf("expected size to be equal to %d, got %d", i, qwu.Sizes[i])
		}
	}
	if uf.Count() != n {
		t.Fatalf("expected count to be equal to n: %d, got: %d\n", n, uf.Count())
	}
}

func TestQuickWeightedUnionTwoElements(t *testing.T) {
	expectedIDs := []int{0, 0}
	uf := CreateQuickWeightedUnion(2)
	uf.Union(0, 1)
	if uf.Count() != 1 {
		t.Fatalf("expected count to be 1, got: %d\n", uf.Count())
	}
	if !uf.Connected(0, 1) {
		t.Fatalf("expected 0 and 1 to be connected, is not connected")
	}
	if !CompareIDs(expectedIDs, uf.IDs()) {
		t.Fatalf("expected ids to same\nexpected: %+v\ngot:%+v\n'", expectedIDs, uf.IDs())
	}
	qwu, ok := (uf).(*QuickWeightedUnion)
	if !ok {
		t.Fatal("created type is not expected type of QuickWeightedUnion")
	}
	expectedSizes := []int{2, 1}
	if !CompareIDs(expectedSizes, qwu.Sizes) {
		t.Fatalf("expected sizes to same\nexpected: %+v\ngot:%+v\n'", expectedSizes, qwu.Sizes)
	}
}

func TestQuickWeightedUnionFourElements(t *testing.T) {
	expectedIDs := []int{0, 1, 2, 3}
	expectedSizes := []int{1, 1, 1, 1}
	uf := CreateQuickWeightedUnion(4)
	qwu, ok := (uf).(*QuickWeightedUnion)
	if !ok {
		t.Fatal("created type is not expected type of QuickWeightedUnion")
	}
	{
		uf.Union(0, 1)
		if uf.Count() != 3 {
			t.Fatalf("expected count to be 3, got: %d\n", uf.Count())
		}
		expectedIDs = []int{0, 0, 2, 3}
		if !CompareIDs(expectedIDs, expectedIDs) {
			t.Fatalf("expected ids to same\nexpected: %+v\ngot:%+v\n'", expectedIDs, uf.IDs())
		}
		expectedSizes = []int{2, 1, 1, 1}
		if !CompareIDs(expectedSizes, qwu.Sizes) {
			t.Fatalf("expected ids to same\nexpected: %+v\ngot:%+v\n'", expectedIDs, uf.IDs())
		}
	}
	{
		uf.Union(2, 3)
		if uf.Count() != 2 {
			t.Fatalf("expected count to be 2, got: %d\n", uf.Count())
		}
		expectedIDs = []int{0, 0, 2, 2}
		if !CompareIDs(expectedIDs, expectedIDs) {
			t.Fatalf("expected ids to same\nexpected: %+v\ngot:%+v\n'", expectedIDs, uf.IDs())
		}
		expectedSizes = []int{2, 1, 2, 1}
		if !CompareIDs(expectedSizes, qwu.Sizes) {
			t.Fatalf("expected ids to same\nexpected: %+v\ngot:%+v\n'", expectedIDs, uf.IDs())
		}
	}
	{
		uf.Union(1, 3)
		if uf.Count() != 1 {
			t.Fatalf("expected count to be 1, got: %d\n", uf.Count())
		}
		expectedIDs = []int{0, 0, 0, 2}
		if !CompareIDs(expectedIDs, expectedIDs) {
			t.Fatalf("expected ids to same\nexpected: %+v\ngot:%+v\n'", expectedIDs, uf.IDs())
		}
		expectedSizes = []int{4, 1, 2, 1}
		if !CompareIDs(expectedSizes, qwu.Sizes) {
			t.Fatalf("expected ids to same\nexpected: %+v\ngot:%+v\n'", expectedIDs, uf.IDs())
		}
	}
}

func TestQuickWeightedUnionTinyInput(t *testing.T) {
	expectedIDs := []int{6, 2, 6, 4, 4, 6, 6, 2, 4, 4}
	expectedSizes := []int{1, 1, 3, 1, 4, 1, 6, 1, 1, 1}
	uf, err := CreateQuickWeightedUnionFromFile("./data/tinyUF.txt")
	if err != nil {
		t.Fatalf("failed to create quick find\n")
	}
	if uf.Count() != 2 {
		t.Fatalf("expected component size of 2, got %d\n", uf.Count())
	}
	if !CompareIDs(expectedIDs, uf.IDs()) {
		t.Fatalf("expected ids to be same\nexpected: %+v\ngot:%+v\n", expectedIDs, uf.IDs())
	}
	qwu, ok := (uf).(*QuickWeightedUnion)
	if !ok {
		t.Fatal("created type is not expected type of QuickWeightedUnion")
	}
	if !CompareIDs(expectedSizes, qwu.Sizes) {
		t.Fatalf("expected sizes to be same\nexpected: %+v\ngot:%+v\n", expectedSizes, qwu.Sizes)
	}
}
