package union_find

type Percolation struct {
	N                  int
	UF                 UnionFind
	VirtualTopIndex    int
	VirtualBottomIndex int
	OpenSiteIDs        []int
}

func CreatePercolation(n int) *Percolation {
	p := &Percolation{}
	p.Reinitialize(n)
	return p
}

func (p *Percolation) Reinitialize(n int) {
	p.N = n
	// create two extra elements representing the virtual top and bottom
	p.UF = CreateQuickWeightedUnion(n*n + 2)
	p.VirtualTopIndex = 0
	p.VirtualBottomIndex = n*n + 1
	// connect top row to virtual top and bottom row to virtual bottom
	for idx := 0; idx < n; idx++ {
		// TODO(nick): these idx are messed up debug this
		topIdx := idx + 1
		p.UF.Union(p.VirtualTopIndex, topIdx)
		bottomIdx := idx + ((n * n) - 1)
		p.UF.Union(p.VirtualBottomIndex, bottomIdx)
	}
}

func (p *Percolation) Translate2DTo1D(x int, y int) int {
	// offset by 1 because the first element is the virtual top index
	return x + (y * p.N) + 1
}

func (p *Percolation) ValidatePosition(x int, y int) bool {
	if x >= 0 && x < p.N && y >= 0 && y < p.N {
		return true
	}
	return false
}

func (p *Percolation) Percolates() bool {
	if p.UF.Connected(p.VirtualTopIndex, p.VirtualBottomIndex) {
		return true
	}
	return false
}
