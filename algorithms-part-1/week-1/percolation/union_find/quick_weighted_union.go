package union_find

type QuickWeightedUnion struct {
	Elements   []int
	Sizes      []int
	Components int
}

func (wu *QuickWeightedUnion) Union(p int, q int) {
	i := wu.Find(p)
	j := wu.Find(q)
	// make smaller root poin to larger one
	if wu.Sizes[i] < wu.Sizes[j] {
		wu.Elements[i] = j
		wu.Sizes[j] += wu.Sizes[i]
	} else {
		wu.Elements[j] = i
		wu.Sizes[i] += wu.Sizes[j]
	}
	wu.Components--
}

func (wu *QuickWeightedUnion) Find(p int) int {
	for p != wu.Elements[p] {
		p = wu.Elements[p]
	}
	return p
}

func (wu *QuickWeightedUnion) Connected(p int, q int) bool {
	result := wu.Find(p) == wu.Find(q)
	return result
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
		Sizes:      make([]int, n),
		Components: n,
	}
	for i := 0; i < n; i++ {
		weightedUnion.Elements[i] = i
		weightedUnion.Sizes[i] = 1
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
