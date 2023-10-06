package union_find

type QuickWeightedUnionCompress struct {
	Elements   []int
	Sizes      []int
	Components int
}

func (uc *QuickWeightedUnionCompress) Union(p int, q int) {
	i := uc.Find(p)
	j := uc.Find(q)
	if i == j {
		return
	}
	// make smaller root poin to larger one
	if uc.Sizes[i] < uc.Sizes[j] {
		uc.Elements[i] = j
		uc.Sizes[j] += uc.Sizes[i]
	} else {
		uc.Elements[j] = i
		uc.Sizes[i] += uc.Sizes[j]
	}
	uc.Components--
}

func (uc *QuickWeightedUnionCompress) Find(p int) int {
	for p != uc.Elements[p] {
		// path compression
		uc.Elements[p] = uc.Elements[uc.Elements[p]]
		p = uc.Elements[p]
	}
	return p
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
	weightedUnionCompressed := QuickWeightedUnionCompress{
		Elements:   make([]int, n),
		Sizes:      make([]int, n),
		Components: n,
	}
	for i := 0; i < n; i++ {
		weightedUnionCompressed.Elements[i] = i
		weightedUnionCompressed.Sizes[i] = 1
	}
	return &weightedUnionCompressed
}

func CreateQuickWeightedUnionCompressedFromFile(path string) (UnionFind, error) {
	file, err := CreateUnionFile(path)
	if err != nil {
		return nil, err
	}
	uf := CreateQuickWeightedUnionCompressed(file.N)
	for i := 0; i < len(file.Pairs); i++ {
		p := file.Pairs[i][0]
		q := file.Pairs[i][1]
		if uf.Connected(p, q) {
			continue
		}
		uf.Union(p, q)
	}
	return uf, nil
}
