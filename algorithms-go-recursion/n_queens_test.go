package recursion

import (
	"testing"

	"gotest.tools/assert"
)

func TestPlaceQueen1(t *testing.T) {
	board := makeBoardQueens(2)
	success := placeQueen1(board, 0, 0)
	dumpBoardString(board)
	assert.Assert(t, success == false)

	board = makeBoardQueens(3)
	success = placeQueen1(board, 0, 0)
	dumpBoardString(board)
	assert.Assert(t, success == false)

	board = makeBoardQueens(4)
	success = placeQueen1(board, 0, 0)
	dumpBoardString(board)
	assert.Assert(t, success == true)
}
