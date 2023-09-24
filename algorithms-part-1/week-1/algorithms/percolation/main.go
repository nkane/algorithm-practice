package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	DefaultRows    = int32(20)
	DefaultColumns = int32(20)
)

var (
	ScreenWidth  = int32(1280)
	ScreenHeight = int32(720)
	CellWidth    = int32(ScreenWidth / DefaultColumns)
	CellHeight   = int32(ScreenHeight / DefaultRows)
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

func (c *Cell) Draw() {
	if c.Open {
		rl.DrawRectangle(int32(c.X)*CellWidth, int32(c.Y)*CellHeight, CellWidth, CellHeight, rl.White)
	} else {
		rl.DrawRectangle(int32(c.X)*CellWidth, int32(c.Y)*CellHeight, CellWidth, CellHeight, rl.Black)
	}
	rl.DrawRectangleLines(int32(c.X)*CellWidth, int32(c.Y)*CellHeight, CellWidth, CellHeight, rl.Black)
}

func main() {
	w := Window{
		Width:  ScreenWidth,
		Height: ScreenHeight,
		FPS:    60,
	}
	rl.InitWindow(w.Width, w.Height, "percolation simulation")
	rl.SetTargetFPS(w.FPS)

	columns := int(DefaultColumns)
	rows := int(DefaultRows)

	grid := [DefaultColumns][DefaultRows]Cell{}

	for i := 0; i < columns; i++ {
		for j := 0; j < rows; j++ {
			grid[i][j] = Cell{
				X: i,
				Y: j,
			}
		}
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		{
			rl.ClearBackground(rl.RayWhite)
			for i := 0; i < columns; i++ {
				for j := 0; j < rows; j++ {
					grid[i][j].Draw()
				}
			}
		}
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
