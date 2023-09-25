package main

import (
	"log"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	DefaultFPS = int32(60)
)

var (
	ScreenWidth  = int32(1280)
	ScreenHeight = int32(720)
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

type State struct {
	MinN       int
	MaxN       int
	N          int32
	CellWidth  int32
	CellHeight int32
}

type Game struct {
	Window *Window
	State  State
	Grid   [][]Cell
}

func CreateGame(n int32, w *Window) *Game {
	g := Game{
		Window: w,
		State: State{
			MinN: 5,
			N:    n,
		},
	}
	g.InitializeGrid(n)
	return &g
}

func (g *Game) InitializeGrid(n int32) {
	g.State.CellWidth = (g.Window.Width / 2) / n
	g.State.CellHeight = g.Window.Height / n
	g.Grid = make([][]Cell, n)
	for column := 0; column < len(g.Grid); column++ {
		g.Grid[column] = make([]Cell, n)
		for row := 0; row < len(g.Grid[column]); row++ {
			g.Grid[column][row] = Cell{
				X:    column,
				Y:    row,
				Open: true,
			}
		}
	}
}

func (g *Game) Render() {
	// render GUI
	if gui.Spinner(rl.NewRectangle(50, 10, 150, 25), "N Value", &g.State.N, g.State.MinN, g.State.MaxN, false) {
		log.Printf("reinit grid\n")
		g.InitializeGrid(g.State.N)
	}
	// render grid
	for column := 0; column < len(g.Grid); column++ {
		for row := 0; row < len(g.Grid[column]); row++ {
			g.Grid[column][row].Draw(g.State.CellWidth, g.State.CellHeight, g.Window.Width/2)
		}
	}
}

func (c *Cell) Draw(cellWidth int32, cellHeight int32, xOffset int32) {
	if c.Open {
		rl.DrawRectangle(int32(c.X)*cellWidth+xOffset, int32(c.Y)*cellHeight, cellWidth, cellHeight, rl.White)
	} else {
		rl.DrawRectangle(int32(c.X)*cellWidth+xOffset, int32(c.Y)*cellHeight, cellWidth, cellHeight, rl.Black)
	}
	rl.DrawRectangleLines(int32(c.X)*cellWidth+xOffset, int32(c.Y)*cellHeight, cellWidth, cellHeight, rl.Black)
}

func main() {
	w := Window{
		Width:  ScreenWidth,
		Height: ScreenHeight,
		FPS:    DefaultFPS,
	}
	rl.InitWindow(w.Width, w.Height, "percolation simulation")
	rl.SetTargetFPS(w.FPS)

	n := int32(5)
	game := CreateGame(n, &w)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		{
			rl.ClearBackground(rl.RayWhite)
			game.Render()
		}
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
