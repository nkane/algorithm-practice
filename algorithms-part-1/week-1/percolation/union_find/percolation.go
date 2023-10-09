package union_find

type Cell struct {
	X                   int
	Y                   int
	Open                bool
	VirtualTopConnected bool
}

type Percolation struct {
	N             int
	UF            UnionFind
	Grid          [][]Cell
	VirtualTop    int
	VirtualBottom int
}

func (p *Percolation) ConnectAdjactSites(i int, j int) {
	directions := [][]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
	for _, dir := range directions {
		ni := i + dir[0]
		nj := j + dir[1]
		if ni >= 0 && ni < p.N && nj >= 0 && nj < p.N && p.Grid[ni][nj].Open {
			p.UF.Union(i*p.N+j, ni*p.N+nj)
		}
		// check if site is connect to top, translate to 1D coordindate
		cellID := i + (j * p.N)
		if p.UF.Connected(p.VirtualTop, cellID) {
			p.Grid[i][j].VirtualTopConnected = true
		}
	}
}
