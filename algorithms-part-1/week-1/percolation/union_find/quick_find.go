package union_find

type QuickFind struct {
	Elements   []int
	Components int
}

func (qf *QuickFind) Union(p int, q int) {
	// put p and q into the same component
	// find component associated with p
	pID := qf.Find(p)
	// find component associatd with q
	qID := qf.Find(q)
	// if p and q are already in the same
	// component, don't do anything
	if pID == qID {
		return
	}
	// iterate over all elements check to see
	// if the element's value is equal to p
	// if it's equal set it to q's value
	for i := 0; i < len(qf.Elements); i++ {
		if qf.Elements[i] == pID {
			qf.Elements[i] = qID
		}
	}
	// reduce component count
	qf.Components--
}

func (qf *QuickFind) Find(p int) int {
	return qf.Elements[p]
}

func (qf *QuickFind) Connected(p int, q int) bool {
	pFound := qf.Find(p)
	qFound := qf.Find(q)
	return pFound == qFound
}

func (qf *QuickFind) Count() int {
	return qf.Components
}

func (qf *QuickFind) IDs() []int {
	return qf.Elements
}

func CreateQuickFind(n int) UnionFind {
	quickFind := QuickFind{
		Elements:   make([]int, n),
		Components: n,
	}
	for i := 0; i < n; i++ {
		quickFind.Elements[i] = i
	}
	return &quickFind
}
