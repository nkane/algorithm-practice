package recursion

func placeQueens4(board [][]string, numRows int, c int) bool {
	if c == numRows {
		return boardIsLegal(board)
	}
	if !boardIsLegal(board) {
		return false
	}
	for r := 0; r < numRows; r++ {
		board[r][c] = "Q"
		if placeQueens4(board, numRows, c+1) {
			return true
		}
		board[r][c] = "."
	}
	return false
}

/*
func placeQueens4(board [][]string, numRows, c int) bool {
	// If c == numRows, then all columns have a queen.
	if c == numRows {
		// Return true if this board is legal.
		return boardIsLegal(board, numRows)
	} else {
		// If the board is already illegal, bail out.
		if !boardIsLegal(board, numRows) {
			return false
		}

		// Try positions in this column.
		for r := 0; r < numRows; r++ {
			// Try placing a queen at [r, c].
			board[r][c] = "Q"

			// Recursively place queens in the following columns.
			// If we found a legal board, return true.
			if placeQueens4(board, numRows, c+1) {
				return true
			}

			// That didn't work. Undo this try.
			board[r][c] = "."
		}

		// If we get to this point, then we couldn't find
		// a legal board given the queens in the earlier columns.
		return false
	}
}
*/
