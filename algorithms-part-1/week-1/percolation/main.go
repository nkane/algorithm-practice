package main

import (
	"fmt"
	"math/rand"
	"percolation/grid"
	"sync"
	"time"

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

type State struct {
	MinN           int
	MaxN           int
	OpenSites      int
	SpinnerMinN    int
	SpinnerMaxN    int
	SpinnerN       int32
	N              int32
	CursorPosition rl.Vector2
	Percolates     bool
	Lock           sync.RWMutex
}

type DebugVec2 struct {
	ID int
	X  int
	Y  int
}

type DebugState struct {
	ReplayOpenOrder []DebugVec2
}

type Simulation struct {
	Window     *Window
	State      State
	DebugState DebugState
	Grid       *grid.Grid
}

func (s *Simulation) Reinitialize(n int32) {
	s.State = State{}
	s.State.N = n
	s.State.SpinnerN = n
	s.State.MinN = 1
	s.State.MaxN = int(n) * int(n)
	s.State.SpinnerMinN = 1
	s.State.SpinnerMaxN = 20
	offset := rl.Vector2{
		X: float32(s.Window.Width / 2),
		Y: float32(0),
	}
	size := rl.Vector2{
		X: float32((s.Window.Width / 2) / n),
		Y: float32((s.Window.Height / n)),
	}
	s.Grid = grid.CreateGrid(int(n), offset, size)
	s.State.Lock.Lock()
	s.State.OpenSites = 0
	s.State.Lock.Unlock()
	s.DebugState = DebugState{}
}

func (s *Simulation) MonteCarloOpen() {
	go func() {
		if !s.State.Percolates {
			if s.State.OpenSites < s.State.MaxN {
				for {
					x := rand.Intn(int(s.State.N))
					y := rand.Intn(int(s.State.N))
					if len(s.DebugState.ReplayOpenOrder) > 0 {
						x = s.DebugState.ReplayOpenOrder[0].X
						y = s.DebugState.ReplayOpenOrder[0].Y
						s.DebugState.ReplayOpenOrder = s.DebugState.ReplayOpenOrder[1:]
					}
					// TODO(nick): can run this in a go routine until it the system percolates
					if s.Grid.OpenAt(x, y) {
						s.State.Lock.Lock()
						s.State.OpenSites++
						s.State.Lock.Unlock()
						time.Sleep(600 * time.Millisecond)
					}
					if s.State.Percolates {
						break
					}
				}
			}
		}
	}()
}

func (s *Simulation) UpdateAndRender() {
	s.State.CursorPosition = rl.GetMousePosition()
	// render UI components
	if gui.Spinner(rl.NewRectangle(150, 10, 150, 25), "N Value", &s.State.SpinnerN, s.State.SpinnerMinN, s.State.SpinnerMaxN, false) {
		if s.State.SpinnerN != s.State.N {
			s.Reinitialize(s.State.SpinnerN)
		}
	}
	if gui.Button(rl.NewRectangle(150, 50, 150, 25), fmt.Sprintf("Open Sites: %d", s.State.OpenSites)) {
		s.MonteCarloOpen()
	}
	if gui.Button(rl.NewRectangle(150, 75, 150, 25), "Reset") {
		s.Reinitialize(s.State.SpinnerN)
	}
	if s.Grid.Percolation.Percolates() {
		gui.Label(rl.NewRectangle(50, 100, 500, 50), "System Percolations")
		s.State.Percolates = true
	}
	s.Grid.SetCursorIntersect(s.State.CursorPosition)
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		if set := s.Grid.OpenAtCursor(); set {
			s.State.Lock.Lock()
			s.State.OpenSites++
			s.State.Lock.Unlock()
		}
	}
	s.Grid.Draw()
}

func main() {
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
