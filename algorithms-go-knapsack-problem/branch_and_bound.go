package knapsack

func BranchAndBound(items []Item, allowedWeight int) ([]Item, int, int) {
	bestValue := 0
	currentValue := 0
	currentWeight := 0
	remainingValue := SumValues(items, true)
	return DoBranchAndBound(items, allowedWeight, 0, bestValue, currentValue, currentWeight, remainingValue)
}

func DoBranchAndBound(
	items []Item,
	allowedWeight int,
	nextIndex int,
	bestValue int,
	currentValue int,
	currentWeight int,
	remainingValue int,
) ([]Item, int, int) {
	if nextIndex >= len(items) {
		solution := CopyItems(items)
		return solution, currentValue, 1
	}
	if currentValue+remainingValue <= bestValue {
		return nil, currentValue, 1
	}
	var test1Solution []Item
	var test1Value int
	var test1Calls int
	if currentWeight+items[nextIndex].Weight <= allowedWeight {
		items[nextIndex].IsSelected = true
		test1Solution, test1Value, test1Calls = DoBranchAndBound(items, allowedWeight, nextIndex+1, bestValue,
			currentValue+items[nextIndex].Value,
			currentWeight+items[nextIndex].Weight,
			remainingValue-items[nextIndex].Value)
		if test1Value > bestValue {
			bestValue = test1Value
		}
	} else {
		test1Solution = nil
		test1Value = 0
		test1Calls = 1
	}

	var test2Solution []Item
	var test2Value int
	var test2Calls int
	if currentValue+remainingValue-items[nextIndex].Value > bestValue {
		items[nextIndex].IsSelected = false
		test2Solution, test2Value, test2Calls = DoBranchAndBound(items, allowedWeight, nextIndex+1, bestValue,
			currentValue, currentWeight, remainingValue-items[nextIndex].Value)
	} else {
		test2Solution = nil
		test2Value = 0
		test2Calls = 1
	}

	if test1Value >= test2Value {
		return test1Solution, test1Value, test1Calls + test2Calls + 1
	} else {
		return test2Solution, test2Value, test1Calls + test2Calls + 1
	}
}
