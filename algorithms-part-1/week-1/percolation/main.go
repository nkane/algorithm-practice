package main

import (
	"fmt"
	uf "percolation/union_find"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// TODO(nick):
// - visualize quick find
// - visualize quick union
// - visualize weighted quick union
// - visualize weighted quick union compressed
// - visualize percolation
const (
	DefaultFPS = int32(60)
)

var (
	ScreenWidth  = int32(1920)
	ScreenHeight = int32(1080)
)

type Window struct {
	Width  int32
	Height int32
	FPS    int32
}

type Cell struct {
	X    int
	Y    int
	Open bool
}

func (c *Cell) Draw(cellWidth int32, cellHeight int32, xOffset int32, n int) {
	if c.Open {
		rl.DrawRectangle(int32(c.X)*cellWidth+xOffset, int32(c.Y)*cellHeight, cellWidth, cellHeight, rl.White)
	} else {
		rl.DrawRectangle(int32(c.X)*cellWidth+xOffset, int32(c.Y)*cellHeight, cellWidth, cellHeight, rl.Black)
	}
	rl.DrawText(fmt.Sprintf("%d: {x: %d, y: %d}", c.X+(c.Y*n), c.X, c.Y), int32(c.X)*cellWidth+xOffset+5, int32(c.Y)*cellHeight+10, 28, rl.Red)
	rl.DrawRectangleLines(int32(c.X)*cellWidth+xOffset, int32(c.Y)*cellHeight, cellWidth, cellHeight, rl.Gray)
}

type State struct {
	MinN       int
	MaxN       int
	OpenSites  int32
	SpinnerN   int32
	N          int32
	CellWidth  int32
	CellHeight int32
}

type Simulation struct {
	Window *Window
	Grid   [][]Cell
	State  State
}

func (s *Simulation) Reinitialize(n int32) {
	s.State.CellWidth = (s.Window.Width / 2) / n
	s.State.CellHeight = (s.Window.Height / n)
	s.State.N = n
	s.State.SpinnerN = n
	s.State.MaxN = 1
	s.State.MaxN = 20
	s.Grid = make([][]Cell, n)
	for column := 0; column < len(s.Grid); column++ {
		s.Grid[column] = make([]Cell, n)
		for row := 0; row < len(s.Grid[column]); row++ {
			s.Grid[column][row] = Cell{
				X:    column,
				Y:    row,
				Open: false,
			}
		}
	}
}

func (s *Simulation) UpdateAndRender() {
	// render UI components
	if gui.Spinner(rl.NewRectangle(150, 10, 150, 25), "N Value", &s.State.SpinnerN, s.State.MinN, s.State.MaxN, false) {
		if s.State.SpinnerN != s.State.N {
			s.Reinitialize(s.State.N)
		}
	}
	// render grid
	for column := 0; column < len(s.Grid); column++ {
		for row := 0; row < len(s.Grid[column]); row++ {
			s.Grid[column][row].Draw(s.State.CellWidth, s.State.CellHeight, s.Window.Width/2, int(s.State.N))
		}
	}
}

func main() {
	uf.CreateQuickFind(1)
	window := &Window{
		Width:  ScreenWidth,
		Height: ScreenHeight,
		FPS:    DefaultFPS,
	}
	rl.InitWindow(window.Width, window.Height, "percolation simulation")
	rl.SetTargetFPS(window.FPS)

	simulation := &Simulation{
		Window: window,
	}
	simulation.Reinitialize(1)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		{
			rl.ClearBackground(rl.RayWhite)
			simulation.UpdateAndRender()
		}
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
