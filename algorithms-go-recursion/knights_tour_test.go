package recursion

import (
	"testing"
	"time"
)

func TestKnightsTour(t *testing.T) {
	initializeOffsets()
	board := makeBoard(numRows, numCols)
	board[0][0] = 0
	start := time.Now()
	if findTour(board, numRows, numCols, 0, 0, 1) {
		t.Log("Successfully found tour")
	} else {
		t.Log("Could not find tour")
	}
	elapsed := time.Since(start)
	dumpBoard(board)
	t.Logf("%f seconds\n", elapsed.Seconds())
	t.Logf("%d calls\n", numCalls)
}
