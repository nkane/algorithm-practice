package grid

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Cell struct {
	X         int
	Y         int
	Open      bool
	Size      rl.Vector2
	Offset    rl.Vector2
	Position  rl.Vector2
	Rectangle rl.Rectangle
}

func (c *Cell) Draw() {
	color := rl.Black
	if c.Open {
		color = rl.White
	}
	rl.DrawRectangleRec(c.Rectangle, color)
	//rl.DrawText(fmt.Sprintf("%d: {x: %d, y: %d}", c.X+(c.Y*n), c.X, c.Y), int32(c.X)*cellWidth+xOffset+5, int32(c.Y)*cellHeight+10, 28, rl.Red)
	rl.DrawRectangleLinesEx(c.Rectangle, 1.0, rl.LightGray)
}

func (c *Cell) DrawWithBorder() {
	c.Draw()
	rl.DrawRectangleLinesEx(c.Rectangle, 10.0, rl.Yellow)
}

type Grid struct {
	Cells    [][]Cell
	N        int
	Offset   rl.Vector2
	CursorAt *rl.Vector2
}

func (g *Grid) Draw() {
	for y := 0; y < len(g.Cells); y++ {
		for x := 0; x < len(g.Cells[0]); x++ {
			if g.CursorAt != nil && int(g.CursorAt.Y) == y && int(g.CursorAt.X) == x {
				g.Cells[y][x].DrawWithBorder()
			} else {
				g.Cells[y][x].Draw()
			}
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
		}
	}
}

func (g *Grid) CloseAtCursor() {
	if g.CursorAt != nil {
		x := int(g.CursorAt.X)
		y := int(g.CursorAt.Y)
		if g.Cells[y][x].Open {
			g.Cells[y][x].Open = false
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
			position := rl.Vector2{
				X: float32(x)*size.X + offset.X,
				Y: float32(y)*size.Y + offset.Y,
			}
			rectangle := rl.NewRectangle(position.X, position.Y, size.X, size.Y)
			grid.Cells[y][x] = Cell{
				X:         x,
				Y:         y,
				Open:      false,
				Offset:    offset,
				Size:      size,
				Position:  position,
				Rectangle: rectangle,
			}
		}
	}
	return &grid
}
