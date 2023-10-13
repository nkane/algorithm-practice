package union_find

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Cell struct {
	X                   int
	Y                   int
	Open                bool
	VirtualTopConnected bool
}

func (c *Cell) Draw(cellWidth int32, cellHeight int32, xOffset int32, n int) {
	if c.Open {
		if c.VirtualTopConnected {
			rl.DrawRectangle(int32(c.X)*cellWidth+xOffset, int32(c.Y)*cellHeight, cellWidth, cellHeight, rl.SkyBlue)
		} else {
			rl.DrawRectangle(int32(c.X)*cellWidth+xOffset, int32(c.Y)*cellHeight, cellWidth, cellHeight, rl.White)
		}
	} else {
		rl.DrawRectangle(int32(c.X)*cellWidth+xOffset, int32(c.Y)*cellHeight, cellWidth, cellHeight, rl.Black)
	}
	rl.DrawText(fmt.Sprintf("%d: {x: %d, y: %d}", c.X+(c.Y*n), c.X, c.Y), int32(c.X)*cellWidth+xOffset+5, int32(c.Y)*cellHeight+10, 28, rl.Red)
	rl.DrawRectangleLines(int32(c.X)*cellWidth+xOffset, int32(c.Y)*cellHeight, cellWidth, cellHeight, rl.Gray)
}

type Percolation struct {
	N                  int
	UF                 UnionFind
	Grid               [][]Cell
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
	p.Grid = make([][]Cell, n)
	for column := 0; column < len(p.Grid); column++ {
		p.Grid[column] = make([]Cell, n)
		for row := 0; row < len(p.Grid[column]); row++ {
			p.Grid[column][row] = Cell{
				X:    column,
				Y:    row,
				Open: false,
			}
		}
	}
	// create two extra elements representing the virtual top and bottom
	p.UF = CreateQuickWeightedUnion(n*n + 2)
	p.VirtualTopIndex = n * n
	p.VirtualBottomIndex = n*n + 1
	for column := 0; column < len(p.Grid); column++ {
		// connect top row to virtual top index
		// translate 2D coordinate to 1D
		cellID := column + (0 * p.N)
		p.UF.Union(p.VirtualTopIndex, cellID)
		cellID = column + ((p.N - 1) * p.N)
		p.UF.Union(p.VirtualBottomIndex, cellID)
	}
}

func (p *Percolation) Translate2DTo1D(x int, y int) int {
	return x + (y * p.N)
}

func (p *Percolation) ValidatePosition(x int, y int) bool {
	if x >= 0 && x < p.N && y >= 0 && y < p.N {
		return true
	}
	return false
}

func (p *Percolation) ConnectAdjactSites(x int, y int) {
	directions := [][]int{
		{0, -1}, // top
		{0, 1},  // bottom
		{-1, 0}, // left
		{1, 0},  // right
	}
	cellID := x + (y * p.N)
	// check if site is connect to top, translate to 1D coordindate
	if p.UF.Connected(p.VirtualTopIndex, cellID) {
		p.Grid[x][y].VirtualTopConnected = true
	}
	p.OpenSiteIDs = append(p.OpenSiteIDs, cellID)
	for _, dir := range directions {
		nx := x + dir[0]
		ny := y + dir[1]
		if p.ValidatePosition(nx, ny) && p.Grid[nx][ny].Open {
			adjacentCellID := nx + (ny * p.N)
			p.UF.Union(cellID, adjacentCellID)
			if p.UF.Connected(p.VirtualTopIndex, adjacentCellID) {
				p.Grid[nx][ny].VirtualTopConnected = true
				// when this gets flagged as connected to virtual top, we need to check
				// all adjact directions not include the previous nx and ny
				for _, dir := range directions {
					nnx := nx + dir[0]
					nny := ny + dir[1]
					/*
						if nnx == x && nny == y {
							continue
						}
					*/
					if p.ValidatePosition(nnx, nny) {
						if p.Grid[nnx][nny].Open {
							p.Grid[nnx][nny].VirtualTopConnected = true
						}
					}
				}
			}
		}
	}
	if p.UF.Connected(p.VirtualTopIndex, cellID) {
		p.Grid[x][y].VirtualTopConnected = true
	}
}

func (p *Percolation) CheckPercolate() bool {
	result := false
	// special case, if the system is a 1x1
	if len(p.Grid) == 1 && len(p.Grid[0]) == 1 {
		// check if open first, then run connected check
		if p.Grid[0][0].Open {
			if p.UF.Connected(p.VirtualTopIndex, p.VirtualBottomIndex) {
				result = true
			}
		}
	} else {
		if p.UF.Connected(p.VirtualTopIndex, p.VirtualBottomIndex) {
			result = true
		}
	}
	return result
}
