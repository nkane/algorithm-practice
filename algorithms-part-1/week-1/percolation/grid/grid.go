package grid

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Cell struct {
	X      int
	Y      int
	Open   bool
	Size   rl.Vector2
	Offset rl.Vector2
}

func (c *Cell) Draw() {
	position := rl.Vector2{
		X: float32(c.X)*c.Size.X + c.Offset.X,
		Y: float32(c.Y)*c.Size.Y + c.Offset.Y,
	}
	color := rl.Black
	if c.Open {
		color = rl.White
	}
	rec := rl.NewRectangle(position.X, position.Y, c.Size.X, c.Size.Y)
	rl.DrawRectangleRec(rec, color)
	rl.DrawRectangleLinesEx(rec, 10.0, rl.Yellow)
	//rl.DrawText(fmt.Sprintf("%d: {x: %d, y: %d}", c.X+(c.Y*n), c.X, c.Y), int32(c.X)*cellWidth+xOffset+5, int32(c.Y)*cellHeight+10, 28, rl.Red)
}

type Grid struct {
	Cells  [][]Cell
	N      int
	Offset rl.Vector2
}

func (g *Grid) Draw() {
	for y := 0; y < len(g.Cells); y++ {
		for x := 0; x < len(g.Cells[0]); x++ {
			g.Cells[y][x].Draw()
		}
	}
}

func CreateGrid(n int, offset rl.Vector2, size rl.Vector2) *Grid {
	grid := Grid{
		N:      n,
		Offset: offset,
		Cells:  make([][]Cell, n),
	}
	for y := 0; y < len(grid.Cells); y++ {
		grid.Cells[y] = make([]Cell, n)
		for x := 0; x < len(grid.Cells[0]); x++ {
			grid.Cells[y][x] = Cell{
				X:      x,
				Y:      y,
				Open:   false,
				Offset: offset,
				Size:   size,
			}
		}
	}
	return &grid
}
