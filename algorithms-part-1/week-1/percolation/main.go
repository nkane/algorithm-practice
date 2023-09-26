package main

import (
	"fmt"

	uf "percolation/union_find"

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
	ScreenWidth  = int32(1280)
	ScreenHeight = int32(720)
)

func main() {
	fmt.Println("vim-go")
	uf.Test()
	rl.InitWindow(ScreenWidth, ScreenHeight, "percolation simulation")
	rl.SetTargetFPS(DefaultFPS)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		{
			rl.ClearBackground(rl.RayWhite)
		}
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
