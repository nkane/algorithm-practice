package recursion

import (
	"fmt"
	"strings"
)

/*
In the N-queens problem, the goal is to find a way to position N queens on an NxN chessboard so that none of them
can attack any of the others. The eight queens problem, or eight queens puzzle, is the N-queens problem on a standard
8x8 chessboard.

The main goal of the proejct's story is to figure out how many queens you can fit on a 20x20 chessboard. This milestone
won't quite get us there, but it will get us started so we can find a solution in the next mile stone.
*/

func makeBoardQueens(numRows int) [][]string {
	numCols := numRows
	board := make([][]string, numRows)
	for r := range board {
		board[r] = make([]string, numCols)
		for c := 0; c < numCols; c++ {
			board[r][c] = "."
		}
	}
	return board
}

func dumpBoardString(board [][]string) {
	sb := strings.Builder{}
	for r := 0; r < len(board); r++ {
		row := board[r]
		for c := 0; c < len(row); c++ {
			sb.WriteString(fmt.Sprintf("%s ", board[r][c]))
		}
		sb.WriteRune('\n')
	}
	fmt.Println(sb.String())
}

func seriesIsLegal(board [][]string, r0 int, c0 int, dr int, dc int) bool {
	numRows := len(board)
	numCols := numRows
	hasQueen := false
	r := r0
	c := c0
	for {
		if board[r][c] == "Q" {
			// if we already have a queen on this row, then this board is not legal
			if hasQueen {
				return false
			}
			// remember that we have a queen on this row
			hasQueen = true
		}
		// move to the next square in the series
		r += dr
		c += dc
		// if we call off the board, then the series is legal
		if r >= numRows || c >= numCols || r < 0 || c < 0 {
			return true
		}
	}
}

func boardIsLegal(board [][]string) bool {
	numRows := len(board)
	// see if each row is legal
	for r := 0; r < numRows; r++ {
		if !seriesIsLegal(board, r, 0, 0, 1) {
			return false
		}
	}
	// see if each column is legal
	for c := 0; c < numRows; c++ {
		if !seriesIsLegal(board, 0, c, 1, 0) {
			return false
		}
	}
	// see if diagonals down to the right are legal
	for r := 0; r < numRows; r++ {
		if !seriesIsLegal(board, r, 0, 1, 1) {
			return false
		}
	}
	for c := 0; c < numRows; c++ {
		if !seriesIsLegal(board, 0, c, 1, 1) {
			return false
		}
	}
	// see if diagonals down to the left are legal
	for r := 0; r < numRows; r++ {
		if !seriesIsLegal(board, r, numRows-1, 1, -1) {
			return false
		}
	}
	for c := 0; c < numRows; c++ {
		if !seriesIsLegal(board, 0, c, 1, -1) {
			return false
		}
	}
	// if we survived this long, then the board is legal
	return true
}

func boardIsASolution(board [][]string) bool {
	if !boardIsLegal(board) {
		return false
	}
	numRows := len(board)
	numQueens := 0
	for r := 0; r < numRows; r++ {
		for c := 0; c < numRows; c++ {
			if board[r][c] == "Q" {
				numQueens += 1
			}
		}
	}
	return numQueens == numRows
}

func placeQueen1(board [][]string, r int, c int) bool {
	numRows := len(board)
	if r >= numRows {
		return boardIsASolution(board)
	}

	nextR := r
	nextC := c + 1
	if nextC >= numRows {
		nextR += 1
		nextC = 0
	}
	if placeQueen1(board, nextR, nextC) {
		return true
	}
	board[r][c] = "Q"
	if placeQueen1(board, nextR, nextC) {
		return true
	}
	board[r][c] = "."
	return false
}
