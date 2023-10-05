package union_find

type QuickWeightedUnionCompress struct {
	Elements   []int
	Sizes      []int
	Components int
}

func (uc *QuickWeightedUnionCompress) Union(p int, q int) {
	// TODO(nick): implement
}

func (uc *QuickWeightedUnionCompress) Find(p int) int {
	// TODO(nick): implement
	return -1
}

func (uc *QuickWeightedUnionCompress) Connected(p int, q int) bool {
	result := uc.Find(p) == uc.Find(q)
	return result
}

func (uc *QuickWeightedUnionCompress) Count() int {
	return uc.Components
}

func (uc *QuickWeightedUnionCompress) IDs() []int {
	return uc.Elements
}

func CreateQuickWeightedUnionCompressed(n int) UnionFind {
	// TODO(nick): implement
	return nil
}

func CreateQuickWeightedUnionCompressedFromFile(path string) (UnionFind, error) {
	// TODO(nick): implement
	return nil, nil
}
