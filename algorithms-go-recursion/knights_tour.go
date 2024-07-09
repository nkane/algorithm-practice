package recursion

import (
	"fmt"
	"strings"
)

/*
In the knight's tour problem, you need to find a path that a knight could use to visit every square on a chessboard
without visiting any square twice. In the open version of the problem, the knight can start and end on any square. In
the closed version, the knight must finish so it could move back to its starting position in one or more move.
*/

/*
Backtracking isn't really something that you do intentionally during this kind of search. It just sort of happens as
a result of the way the recursive function call works. When one recursive call hits a dead end, control passes back
up to the code that called it and that code tries again with another next move. That's backtracking.

For example, supposed the findTour function adds move 1 to the path and calls itself. The second call to findTour
adds move 2 to the path and calls itself again. The third call adds move 3 to the path, but move 3 is off the board,
so it has hit a dead end.

At that point, the third call to findTour returns false to tell the second call that it could not find a solution.
The second call then tries again with a new move 3b.

This is what we mean by backtracking. The program backtracked up to the second call and tried again with a new move
from that position.

If the new third call cannot find a solution after move 3b, then it will backtrack again and the second call will
try another move 3c.

If the second call to findTour tries all of the possible moves after move 2 and none of them work, then the program
backtracks up even higher to the first call to findTour and tries a new move 2b.

The whole process repeats with the program racing up and down the call tree trying moves and backtracking when it
hits dead ends until it either finds a workable solution or has exhausted every possible combination of moves.
*/

const (
	numRows            = 8
	numCols            = numRows
	requiredClosedTour = false
	unvisited          = -1
)

type Offset struct {
	DR int
	DC int
}

var (
	moveOffsets []Offset
	numCalls    int64
)

func initializeOffsets() {
	moveOffsets = []Offset{
		{-2, -1},
		{-1, -2},
		{+2, -1},
		{+1, -2},
		{-2, +1},
		{-1, +2},
		{+2, +1},
		{+1, +2},
	}
}

func makeBoard(numRows int, numCols int) [][]int {
	board := make([][]int, numRows)
	for r := 0; r < numRows; r++ {
		board[r] = make([]int, numCols)
		for c := 0; c < numCols; c++ {
			board[r][c] = unvisited
		}
	}
	return board
}

func dumpBoard(board [][]int) {
	sb := strings.Builder{}
	for r := 0; r < len(board); r++ {
		row := board[r]
		for c := 0; c < len(row); c++ {
			sb.WriteString(fmt.Sprintf("%02d ", board[r][c]))
		}
		sb.WriteRune('\n')
	}
	fmt.Println(sb.String())
}

func findTour(board [][]int, numRows int, numCols int, curRow int, curCol int, numVisited int) bool {
	numCalls++
	if numVisited == numRows*numCols {
		if requiredClosedTour == false {
			return true
		}
		// check to see whether this is a closed tour
		// loop through the allowed moves from this position, if any of those moves land on a spot where
		// board[r][c] is 0, then we can move back to the starting position, that would be a closed
		// tour, and the function should return true
		for _, offset := range moveOffsets {
			r := curRow + offset.DR
			c := curCol + offset.DC
			if board[r][c] != 0 {
				return true
			}
		}
	}
	// knight has not visited every square
	for _, offset := range moveOffsets {
		// get the move
		r := curRow + offset.DR
		c := curCol + offset.DC
		// see if this move is on the board
		if r < 0 || r >= numRows {
			continue
		}
		if c < 0 || c >= numCols {
			continue
		}
		// see if we have already visisted this position
		if board[r][c] != unvisited {
			continue
		}
		// the move to [r][c] is viable, make this move
		board[r][c] = numVisited
		// try to find the rest of the tour
		if findTour(board, numRows, numCols, r, c, numVisited+1) {
			return true
		}
		// didn't find a tour with this move, unmake this move
		board[r][c] = unvisited
	}
	return false
}
