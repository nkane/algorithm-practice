package recursion

import (
	"testing"

	"gotest.tools/assert"
)

func TestPlaceQueens4(t *testing.T) {
	board := makeBoardQueens(2)
	success := placeQueens4(board, len(board), 0)
	dumpBoardString(board)
	assert.Assert(t, success == false)

	board = makeBoardQueens(4)
	success = placeQueens4(board, len(board), 0)
	dumpBoardString(board)
	assert.Assert(t, success == true)

	board = makeBoardQueens(20)
	success = placeQueens4(board, len(board), 0)
	dumpBoardString(board)
	assert.Assert(t, success == true)
}
