package union_find

type QuickUnion struct {
	Elements   []int
	Components int
}

func (qu *QuickUnion) Union(p int, q int) {
	pID := qu.Find(p)
	qID := qu.Find(q)
	if pID == qID {
		return
	}
	qu.Elements[pID] = qID
	qu.Components--
}

func (qu *QuickUnion) Find(p int) int {
	// find component
	for p != qu.Elements[p] {
		p = qu.Elements[p]
	}
	return p
}

func (qu *QuickUnion) Connected(p int, q int) bool {
	pID := qu.Find(p)
	qID := qu.Find(q)
	return pID == qID
}

func (qu *QuickUnion) Count() int {
	return qu.Components
}

func (qu *QuickUnion) IDs() []int {
	return qu.Elements
}

func CreateQuickUnion(n int) UnionFind {
	quickUnion := QuickUnion{
		Elements:   make([]int, n),
		Components: n,
	}
	for i := 0; i < n; i++ {
		quickUnion.Elements[i] = i
	}
	return &quickUnion
}

func CreateQuickUnionFromFile(path string) (UnionFind, error) {
	file, err := CreateUnionFile(path)
	if err != nil {
		return nil, err
	}
	uf := CreateQuickUnion(file.N)
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
