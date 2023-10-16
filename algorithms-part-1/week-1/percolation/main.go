package main

import (
	"fmt"

	"percolation/grid"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

/*
TODO(nick):
- add ability to open cells by click instead of relying on random order
- this percolation algorithm isn't working as expected, need to debug and look into what is going on
- blog: https://coderanch.com/t/604648/java/Code-review-Percolation-java
- maybe merge both grid and percolation?
*/

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
	//Percolation *uf.Percolation
	Grid *grid.Grid
}

func (s *Simulation) Reinitialize(n int32) {
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
	s.State.OpenSites = 0
	//s.Percolation = uf.CreatePercolation(int(n))

	// NOTE(nick): debug state
	s.DebugState = DebugState{}
	/*
		debug2x2_case_1 := false
		if debug2x2_case_1 {
			// order: 1 -> 0
			s.DebugState.ReplayOpenOrder = append(s.DebugState.ReplayOpenOrder, DebugVec2{
				ID: 1,
				X:  1,
				Y:  0,
			})
			s.DebugState.ReplayOpenOrder = append(s.DebugState.ReplayOpenOrder, DebugVec2{
				ID: 0,
				X:  0,
				Y:  0,
			})
		}
		debug2x2_case_2 := false
		if debug2x2_case_2 {
			// order: 3 -> 0 -> 1 -> 2
			s.DebugState.ReplayOpenOrder = append(s.DebugState.ReplayOpenOrder, DebugVec2{
				ID: 3,
				X:  1,
				Y:  1,
			})
			s.DebugState.ReplayOpenOrder = append(s.DebugState.ReplayOpenOrder, DebugVec2{
				ID: 0,
				X:  0,
				Y:  0,
			})
			s.DebugState.ReplayOpenOrder = append(s.DebugState.ReplayOpenOrder, DebugVec2{
				ID: 1,
				X:  1,
				Y:  0,
			})
			s.DebugState.ReplayOpenOrder = append(s.DebugState.ReplayOpenOrder, DebugVec2{
				ID: 2,
				X:  0,
				Y:  1,
			})
		}
		debug2x2_case_3 := false
		if debug2x2_case_3 {
			// order: 3 -> 0 -> 1
			s.DebugState.ReplayOpenOrder = []DebugVec2{
				{
					ID: 3,
					X:  1,
					Y:  1,
				},
				{
					ID: 0,
					X:  0,
					Y:  0,
				},
				{
					ID: 1,
					X:  1,
					Y:  0,
				},
			}
		}
		debug2x2_case_4 := false
		if debug2x2_case_4 {
			// order: 3 -> 0 -> 2
			s.DebugState.ReplayOpenOrder = []DebugVec2{
				{
					ID: 3,
					X:  1,
					Y:  1,
				},
				{
					ID: 0,
					X:  0,
					Y:  0,
				},
				{
					ID: 2,
					X:  0,
					Y:  1,
				},
			}
		}
		debug2x2_case_5 := false
		if debug2x2_case_5 {
			// order: 2 -> 3 -> 0
			s.DebugState.ReplayOpenOrder = []DebugVec2{
				{
					ID: 2,
					X:  0,
					Y:  1,
				},
				{
					ID: 3,
					X:  1,
					Y:  1,
				},
				{
					ID: 0,
					X:  0,
					Y:  0,
				},
			}
		}
		debug2x2_case_6 := false
		if debug2x2_case_6 {
			// order: 1 -> 0 -> 2
			s.DebugState.ReplayOpenOrder = []DebugVec2{
				{
					ID: 1,
					X:  1,
					Y:  0,
				},
				{
					ID: 0,
					X:  0,
					Y:  0,
				},
				{
					ID: 2,
					X:  0,
					Y:  1,
				},
			}
		}
		debug2x2_case_7 := true
		if debug2x2_case_7 {
			// order: 1 -> 3
			s.DebugState.ReplayOpenOrder = []DebugVec2{
				{
					ID: 1,
					X:  1,
					Y:  0,
				},
				{
					ID: 3,
					X:  1,
					Y:  1,
				},
			}
		}
	*/
}

func (s *Simulation) MonteCarloOpen() {
	/*
		if s.State.OpenSites < s.State.MaxN {
			s.State.OpenSites++
			for {
				x := rand.Intn(int(s.State.N))
				y := rand.Intn(int(s.State.N))
				if len(s.DebugState.ReplayOpenOrder) > 0 {
					x = s.DebugState.ReplayOpenOrder[0].X
					y = s.DebugState.ReplayOpenOrder[0].Y
					s.DebugState.ReplayOpenOrder = s.DebugState.ReplayOpenOrder[1:]
				}
				if !s.Percolation.Grid[x][y].Open {
					s.Percolation.Grid[x][y].Open = true
					s.Percolation.ConnectAdjactSites(x, y)
					break
				}
			}
		}
	*/
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

	// TODO(nick): uncomment this
	/*
		if s.Percolation.CheckPercolate() {
			gui.Label(rl.NewRectangle(50, 100, 500, 50), "System Percolations")
		}
	*/

	s.Grid.SetCursorIntersect(s.State.CursorPosition)
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		s.Grid.OpenAtCursor()
	} else if rl.IsMouseButtonPressed(rl.MouseRightButton) {
		//s.Grid.CloseAtCursor()
	}

	s.Grid.Draw()

	// TODO(nick): old code, remove this
	/*
		// render grid
		for column := 0; column < len(s.Percolation.Grid); column++ {
			for row := 0; row < len(s.Percolation.Grid[column]); row++ {
				s.Percolation.Grid[column][row].Draw(s.State.CellWidth, s.State.CellHeight, s.Window.Width/2, int(s.State.N))
			}
		}
	*/
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
