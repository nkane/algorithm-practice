package recursion

import (
	"testing"

	"gotest.tools/assert"
)

func TestPlaceQueen1(t *testing.T) {
	board := makeBoardQueens(2)
	success := placeQueens1(board, 0, 0)
	dumpBoardString(board)
	assert.Assert(t, success == false)

	board = makeBoardQueens(3)
	success = placeQueens1(board, 0, 0)
	dumpBoardString(board)
	assert.Assert(t, success == false)

	board = makeBoardQueens(4)
	success = placeQueens1(board, 0, 0)
	dumpBoardString(board)
	assert.Assert(t, success == true)
}

func TestPlaceQueen2(t *testing.T) {
	board := makeBoardQueens(2)
	success := placeQueens2(board, 0, 0, 0)
	dumpBoardString(board)
	assert.Assert(t, success == false)

	board = makeBoardQueens(3)
	success = placeQueens2(board, 0, 0, 0)
	dumpBoardString(board)
	assert.Assert(t, success == false)

	board = makeBoardQueens(4)
	success = placeQueens2(board, 0, 0, 0)
	dumpBoardString(board)
	assert.Assert(t, success == true)
}

func TestPlaceQueen3(t *testing.T) {
	board := makeBoardQueens(2)
	success := placeQueens3(board, 0)
	dumpBoardString(board)
	assert.Assert(t, success == false)

	board = makeBoardQueens(3)
	success = placeQueens3(board, 0)
	dumpBoardString(board)
	assert.Assert(t, success == false)

	board = makeBoardQueens(4)
	success = placeQueens3(board, 0)
	dumpBoardString(board)
	assert.Assert(t, success == true)
}
