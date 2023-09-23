package algorithms

import "testing"

func TestQuickFindTiny(t *testing.T) {
	expectedOutput := []int{1, 1, 1, 8, 8, 1, 1, 1, 8, 8}
	uf, err := CreateUnionFindFromFile("./data/tinyUF.txt", QuickFindType)
	if err != nil {
		t.Fatalf("expected error to be nil, received error: %+v", err)
	}
	if uf.GetCount() != 2 {
		t.Fatalf("expected initial count to be 2, got %d", uf.GetCount())
	}
	ids := uf.GetIDs()
	if len(ids) != len(expectedOutput) {
		t.Fatalf("expected lengths do not match, got %d", len(ids))
	}
	for i := 0; i < len(ids); i++ {
		if ids[i] != expectedOutput[i] {
			t.Fatalf("failed comparison check at index %d", i)
		}
	}
}

func TestQuickUnionFindTiny(t *testing.T) {
	expectedOutput := []int{1, 1, 1, 8, 3, 0, 5, 1, 8, 8}
	uf, err := CreateUnionFindFromFile("./data/tinyUF.txt", QuickUnionFindType)
	if err != nil {
		t.Fatalf("expected error to be nil, received error: %+v", err)
	}
	if uf.GetCount() != 2 {
		t.Fatalf("expected initial count to be 2, got %d", uf.GetCount())
	}
	ids := uf.GetIDs()
	if len(ids) != len(expectedOutput) {
		t.Fatalf("expected lengths do not match, got %d", len(ids))
	}
	for i := 0; i < len(ids); i++ {
		if ids[i] != expectedOutput[i] {
			t.Fatalf("failed comparison check at index %d", i)
		}
	}
}

func TestWeightedQuickUnionFind(t *testing.T) {
	expectedOutput := []int{6, 2, 6, 4, 4, 6, 6, 2, 4, 4}
	uf, err := CreateUnionFindFromFile("./data/tinyUF.txt", WeightedQuickUnionFindType)
	if err != nil {
		t.Fatalf("expected error to be nil, received error: %+v", err)
	}
	if uf.GetCount() != 2 {
		t.Fatalf("expected initial count to be 2, got %d", uf.GetCount())
	}
	ids := uf.GetIDs()
	if len(ids) != len(expectedOutput) {
		t.Fatalf("expected lengths do not match, got %d", len(ids))
	}
	for i := 0; i < len(ids); i++ {
		if ids[i] != expectedOutput[i] {
			t.Fatalf("failed comparison check at index %d", i)
		}
	}
}

func TestQuickFindCompressionTiny(t *testing.T) {
	uf := CreateQuickFindPathCompression(10)
	uf.Union(1, 2)
	uf.Union(1, 3)
}
