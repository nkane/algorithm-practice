package union_find

import (
	"math/rand"
	"testing"
)

func TestQuickWeightedUnionCompressedCreate(t *testing.T) {
	n := rand.Intn(1024)
	uf := CreateQuickWeightedUnionCompressed(n)
	cu, ok := (uf).(*QuickWeightedUnionCompress)
	if !ok {
		t.Fatalf("created type is not expected type of QuickWeightedUnionCompress")
	}
	if len(cu.Elements) != len(cu.Sizes) {
		t.Fatal("expected elements and size slices to be the same length")
	}
	for i := 0; i < len(cu.Elements); i++ {
		if cu.Elements[i] != i {
			t.Fatalf("expected element to be equal to %d, got %d", i, cu.Elements[i])
		}
		if cu.Sizes[i] != 1 {
			t.Fatalf("expected size to be equal to %d, got %d", i, cu.Sizes[i])
		}
	}
	if uf.Count() != n {
		t.Fatalf("expected count to be equal to n: %d, got: %d\n", n, uf.Count())
	}
}

func TestQuickWeightedUnionCompressedTwoElements(t *testing.T) {
	expectedIDs := []int{0, 0}
	uf := CreateQuickWeightedUnionCompressed(2)
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
	cu, ok := (uf).(*QuickWeightedUnionCompress)
	if !ok {
		t.Fatal("created type is not expected type of QuickWeightedUnionCompress")
	}
	expectedSizes := []int{2, 1}
	if !CompareIDs(expectedSizes, cu.Sizes) {
		t.Fatalf("expected sizes to same\nexpected: %+v\ngot:%+v\n'", expectedSizes, cu.Sizes)
	}
}

func TestQuickWeightedUnionCompressedFourElements(t *testing.T) {
	expectedIDs := []int{0, 1, 2, 3}
	uf := CreateQuickWeightedUnionCompressed(4)
	cu, ok := (uf).(*QuickWeightedUnionCompress)
	if !ok {
		t.Fatal("created type is not expected type of QuickWeightedUnionCompress")
	}
	{
		uf.Union(0, 1)
		if uf.Count() != 3 {
			t.Fatalf("expected count to be 3, got: %d\n", uf.Count())
		}
		if !uf.Connected(0, 1) {
			t.Fatalf("expected 0 and 1 to be connected, is not connected")
		}
		expectedIDs = []int{0, 0, 2, 3}
		if !CompareIDs(expectedIDs, uf.IDs()) {
			t.Fatalf("expected ids to same\nexpected: %+v\ngot:%+v\n'", expectedIDs, uf.IDs())
		}
		expectedSizes := []int{2, 1, 1, 1}
		if !CompareIDs(expectedSizes, cu.Sizes) {
			t.Fatalf("expected sizes to same\nexpected: %+v\ngot:%+v\n'", expectedSizes, cu.Sizes)
		}
	}
	{
		uf.Union(2, 3)
		if uf.Count() != 2 {
			t.Fatalf("expected count to be 2, got: %d\n", uf.Count())
		}
		if !uf.Connected(2, 3) {
			t.Fatalf("expected 2 and 3 to be connected, is not connected")
		}
		expectedIDs = []int{0, 0, 2, 2}
		if !CompareIDs(expectedIDs, uf.IDs()) {
			t.Fatalf("expected ids to same\nexpected: %+v\ngot:%+v\n'", expectedIDs, uf.IDs())
		}
		expectedSizes := []int{2, 1, 2, 1}
		if !CompareIDs(expectedSizes, cu.Sizes) {
			t.Fatalf("expected sizes to same\nexpected: %+v\ngot:%+v\n'", expectedSizes, cu.Sizes)
		}
	}
	{
		uf.Union(1, 3)
		if uf.Count() != 1 {
			t.Fatalf("expected count to be 2, got: %d\n", uf.Count())
		}
		expectedIDs = []int{0, 0, 0, 2}
		if !CompareIDs(expectedIDs, uf.IDs()) {
			t.Fatalf("expected ids to same\nexpected: %+v\ngot:%+v\n'", expectedIDs, uf.IDs())
		}
		// check if connect, runs path compression
		if !uf.Connected(1, 3) {
			t.Fatalf("expected 1 and 3 to be connected, is not connected")
		}
		expectedIDs = []int{0, 0, 0, 0}
		if !CompareIDs(expectedIDs, uf.IDs()) {
			t.Fatalf("expected ids to same\nexpected: %+v\ngot:%+v\n'", expectedIDs, uf.IDs())
		}
		expectedSizes := []int{4, 1, 2, 1}
		if !CompareIDs(expectedSizes, cu.Sizes) {
			t.Fatalf("expected sizes to same\nexpected: %+v\ngot:%+v\n'", expectedSizes, cu.Sizes)
		}
	}
}

func TestQuickWeightedUnionCompressedTinyInput(t *testing.T) {
	expectedIDs := []int{6, 6, 6, 4, 4, 6, 6, 6, 4, 4}
	expectedSizes := []int{1, 1, 3, 1, 4, 1, 6, 1, 1, 1}
	uf, err := CreateQuickWeightedUnionCompressedFromFile("./data/tinyUF.txt")
	if err != nil {
		t.Fatalf("failed to create quick find\n")
	}
	if uf.Count() != 2 {
		t.Fatalf("expected component size of 2, got %d\n", uf.Count())
	}
	if !CompareIDs(expectedIDs, uf.IDs()) {
		t.Fatalf("expected ids to be same\nexpected: %+v\ngot:%+v\n", expectedIDs, uf.IDs())
	}
	qwu, ok := (uf).(*QuickWeightedUnionCompress)
	if !ok {
		t.Fatal("created type is not expected type of QuickWeightedUnion")
	}
	if !CompareIDs(expectedSizes, qwu.Sizes) {
		t.Fatalf("expected sizes to be same\nexpected: %+v\ngot:%+v\n", expectedSizes, qwu.Sizes)
	}
}
