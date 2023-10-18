package grid

import (
	uf "percolation/union_find"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Cell struct {
	ID                  int
	X                   int
	Y                   int
	Open                bool
	VirtualTopConnected bool
	Size                rl.Vector2
	Offset              rl.Vector2
	Position            rl.Vector2
	Rectangle           rl.Rectangle
}

func (c *Cell) Draw(hoverCell bool, connectedToVirtualTop bool) {
	color := rl.Black
	if c.Open {
		if connectedToVirtualTop {
			color = rl.SkyBlue
		} else {
			color = rl.White
		}
	}
	rl.DrawRectangleRec(c.Rectangle, color)
	//rl.DrawText(fmt.Sprintf("%d: {x: %d, y: %d}", c.X+(c.Y*n), c.X, c.Y), int32(c.X)*cellWidth+xOffset+5, int32(c.Y)*cellHeight+10, 28, rl.Red)
	rl.DrawRectangleLinesEx(c.Rectangle, 1.0, rl.LightGray)
	if hoverCell {
		rl.DrawRectangleLinesEx(c.Rectangle, 10.0, rl.Yellow)
	}
}

type Grid struct {
	Cells       [][]Cell
	N           int
	Offset      rl.Vector2
	CursorAt    *rl.Vector2
	Percolation *uf.Percolation
}

func (g *Grid) Draw() {
	for y := 0; y < len(g.Cells); y++ {
		for x := 0; x < len(g.Cells[0]); x++ {
			idx := g.Percolation.Translate2DTo1D(x, y)
			// check if cell is connected to top
			connectedToVirtualTop := g.Percolation.UF.Connected(g.Percolation.VirtualTopIndex, idx)
			hoverCell := g.CursorAt != nil && int(g.CursorAt.Y) == y && int(g.CursorAt.X) == x
			g.Cells[y][x].Draw(hoverCell, connectedToVirtualTop)
		}
	}
}

func (g *Grid) SetCursorIntersect(cursorPosition rl.Vector2) {
	g.CursorAt = nil
	for y := 0; y < len(g.Cells); y++ {
		for x := 0; x < len(g.Cells[0]); x++ {
			if rl.CheckCollisionPointRec(cursorPosition, g.Cells[y][x].Rectangle) {
				g.CursorAt = &rl.Vector2{
					X: float32(x),
					Y: float32(y),
				}
				return
			}
		}
	}
}

func (g *Grid) OpenAtCursor() {
	if g.CursorAt != nil {
		x := int(g.CursorAt.X)
		y := int(g.CursorAt.Y)
		if !g.Cells[y][x].Open {
			g.Cells[y][x].Open = true
			idx := g.Percolation.Translate2DTo1D(x, y)
			if y == 0 {
				g.Percolation.UF.Union(g.Percolation.VirtualTopIndex, idx)
			}
			if y == g.N-1 {
				g.Percolation.UF.Union(g.Percolation.VirtualBottomIndex, idx)
			}
			if g.Percolation.UF.Connected(g.Percolation.VirtualTopIndex, idx) {
				g.Cells[y][x].VirtualTopConnected = true
			}
			directions := [][]int{
				{1, 0},  // right
				{-1, 0}, // left
				{0, 1},  // down
				{0, -1}, // up
			}
			for _, dir := range directions {
				nX := x + dir[0]
				nY := y + dir[1]
				g.OpenAt(x, y, nX, nY)
			}
		}
	}
}

func (g *Grid) OpenAt(x int, y int, nX int, nY int) {
	if g.Percolation.ValidatePosition(nX, nY) {
		if g.Cells[nY][nX].Open {
			idx := g.Percolation.Translate2DTo1D(x, y)
			nIdx := g.Percolation.Translate2DTo1D(nX, nY)
			g.Percolation.UF.Union(idx, nIdx)
			if g.Percolation.UF.Connected(g.Percolation.VirtualTopIndex, nIdx) {
				g.Cells[y][x].VirtualTopConnected = true
				g.Cells[nY][nX].VirtualTopConnected = true
				// open any adjacent cells
			}
		}
	}
}

func CreateGrid(n int, offset rl.Vector2, size rl.Vector2) *Grid {
	grid := Grid{
		N:           n,
		Offset:      offset,
		Cells:       make([][]Cell, n),
		Percolation: uf.CreatePercolation(int(n)),
	}
	id := 1
	for y := 0; y < n; y++ {
		grid.Cells[y] = make([]Cell, n)
		for x := 0; x < n; x++ {
			position := rl.Vector2{
				X: float32(x)*size.X + offset.X,
				Y: float32(y)*size.Y + offset.Y,
			}
			rectangle := rl.NewRectangle(position.X, position.Y, size.X, size.Y)
			grid.Cells[y][x] = Cell{
				ID:        id,
				X:         x,
				Y:         y,
				Open:      false,
				Offset:    offset,
				Size:      size,
				Position:  position,
				Rectangle: rectangle,
			}
			id++
		}
	}
	return &grid
}
