package union_find

type QuickWeightedUnion struct {
	Elements   []int
	Components int
}

func (wu *QuickWeightedUnion) Union(p int, q int) {
	// TODO(nick): implement
}

func (wu *QuickWeightedUnion) Find(p int) int {
	// TODO(nick): implement
	return -1
}

func (wu *QuickWeightedUnion) Connected(p int, q int) bool {
	// TODO(nick): implement
	return false
}

func (qw *QuickWeightedUnion) Count() int {
	return qw.Components
}

func (qw *QuickWeightedUnion) IDs() []int {
	return qw.Elements
}

func CreateQuickWeightedUnion(n int) UnionFind {
	weightedUnion := QuickWeightedUnion{
		Elements:   make([]int, n),
		Components: n,
	}
	for i := 0; i < n; i++ {
		weightedUnion.Elements[i] = i
	}
	return &weightedUnion
}

func CreateQuickWeightedUnionFromFile(path string) (UnionFind, error) {
	file, err := CreateUnionFile(path)
	if err != nil {
		return nil, err
	}
	uf := CreateQuickWeightedUnion(file.N)
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
