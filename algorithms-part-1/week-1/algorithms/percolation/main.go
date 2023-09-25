package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"

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
	MinN                int
	MaxN                int
	OpenSites           int32
	N                   int32
	CellWidth           int32
	CellHeight          int32
	Files               []fs.DirEntry
	FileNames           []string
	FileListFocus       int32
	FileListScrollIndex int32
	FileListActiveIndex int32
}

type Game struct {
	Window *Window
	State  State
	Grid   [][]Cell
}

func CreateGame(n int32, w *Window) (*Game, error) {
	g := Game{
		Window: w,
	}
	g.ResetGame(n)
	return &g, nil
}

func (g *Game) ResetGame(n int32) {
	g.ResetState(n)
	g.InitializeGrid(n)
}

func (g *Game) ResetState(n int32) error {
	g.State = State{
		MinN:      5,
		MaxN:      20,
		N:         n,
		OpenSites: 0,
	}
	var err error
	g.State.Files, err = os.ReadDir("./grid-data")
	if err != nil {
		return err
	}
	for _, f := range g.State.Files {
		g.State.FileNames = append(g.State.FileNames, f.Name())
	}
	return nil
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
				Open: false,
			}
		}
	}
}

func (g *Game) UpdateAndRender() {
	// render GUI
	if gui.Spinner(rl.NewRectangle(150, 10, 150, 25), "N Value", &g.State.N, g.State.MinN, g.State.MaxN, false) {
		g.ResetGame(g.State.N)
	}
	gui.SetStyle(gui.TEXTBOX, gui.TEXT_ALIGNMENT, gui.TEXT_ALIGN_CENTER)
	if gui.Button(rl.NewRectangle(150, 50, 150, 25), fmt.Sprintf("Open Sites: %d", g.State.OpenSites)) {
		if g.State.OpenSites < int32(g.State.MaxN)*int32(g.State.MaxN) {
			g.State.OpenSites++
			for {
				x := rand.Intn(int(g.State.N))
				y := rand.Intn(int(g.State.N))
				if !g.Grid[x][y].Open {
					g.Grid[x][y].Open = true
					return
				}
			}
		}
	}
	g.State.FileListActiveIndex = gui.ListViewEx(rl.NewRectangle(150, 150, 250, 150),
		g.State.FileNames,
		&g.State.FileListFocus,
		&g.State.FileListScrollIndex,
		g.State.FileListActiveIndex)
	if gui.Button(rl.NewRectangle(200, 300, 150, 25), "Open File") {
		// TODO(nick): load grid data first, then create use union find data
		log.Printf("active index: %d", int(g.State.FileListActiveIndex))
		if int(g.State.FileListActiveIndex) < 0 || int(g.State.FileListActiveIndex) > len(g.State.FileNames) {
			goto skip
		}
		fileName := g.State.FileNames[g.State.FileListActiveIndex]
		log.Printf("attempting to open %s", fileName)
		f, err := os.Open(fmt.Sprintf("grid-data/%s", fileName))
		if err != nil {
			goto skip
		}
		fs := bufio.NewScanner(f)
		fs.Split(bufio.ScanLines)
		if !fs.Scan() {
			log.Printf("failed to read first line of grid data file\n")
			goto skip
		}
		nString := fs.Text()
		n, err := strconv.Atoi(nString)
		if err != nil {
			log.Printf("failed to convert first line to int\n")
			goto skip
		}
		g.ResetGame(int32(n))
		row := 0
		for fs.Scan() {
			line := fs.Text()
			tokens := strings.Fields(line)
			log.Printf("token set: %+v", tokens)
			if len(tokens) >= 0 {
				for column := 0; column < len(tokens); column++ {
					switch tokens[column] {
					case "1":
						g.Grid[column][row].Open = true
						g.State.OpenSites++
					case "0":
						fallthrough
					default:
						g.Grid[column][row].Open = false
					}
				}
			}
			row++
		}
	}
skip:

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
	game, err := CreateGame(n, &w)
	if err != nil {
		panic(err)
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		{
			rl.ClearBackground(rl.RayWhite)
			game.UpdateAndRender()
		}
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
