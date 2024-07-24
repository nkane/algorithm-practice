package knapsack

func ExhaustiveSearch(items []Item, allowedWeight int) ([]Item, int, int) {
	return DoExhaustiveSearch(items, allowedWeight, 0)
}

func DoExhaustiveSearch(items []Item, allowedWeight int, nextIndex int) ([]Item, int, int) {
	// see if we have a full assignment
	if nextIndex >= len(items) {
		// make a copy of this assignment
		assignment := CopyItems(items)
		// return the assignment and its total value
		return assignment, SolutionValue(assignment, allowedWeight), 1
	}
	// don't have a full assignment, try to adding the next item
	items[nextIndex].IsSelected = true
	// recursively call the function
	test1Solution, test1Value, test1Calls := DoExhaustiveSearch(items, allowedWeight, nextIndex+1)
	// try not adding the next item
	items[nextIndex].IsSelected = false
	// recursively call the function
	test2Solution, test2Value, test2Calls := DoExhaustiveSearch(items, allowedWeight, nextIndex+1)
	if test1Value >= test2Value {
		return test1Solution, test1Value, test1Calls + test2Calls + 1
	} else {
		return test2Solution, test2Value, test1Calls + test2Calls + 1
	}
}
