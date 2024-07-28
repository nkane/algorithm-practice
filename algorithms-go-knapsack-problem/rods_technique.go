package knapsack

import "sort"

func RodsTechnique(items []Item, allowedWeight int) ([]Item, int, int) {
	MakeBlockLists(items)
	bestValue := 0
	currentValue := 0
	currentWeight := 0
	remainingValue := SumValues(items, true)
	return DoRodsTechnique(items, allowedWeight, 0, bestValue, currentValue, currentWeight, remainingValue)
}

func RodsTechniqueSorted(items []Item, allowedWeight int) ([]Item, int, int) {
	MakeBlockLists(items)
	sort.Slice(items, func(i, j int) bool {
		return len(items[i].BlockList) > len(items[j].BlockList)
	})
	for i := range items {
		items[i].ID = i
	}
	MakeBlockLists(items)
	bestValue := 0
	currentValue := 0
	currentWeight := 0
	remainingValue := SumValues(items, true)
	return DoRodsTechnique(items, allowedWeight, 0,
		bestValue, currentValue, currentWeight, remainingValue)
}

func DoRodsTechnique(
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
	test1Solution = nil
	test1Value := 0
	test1Calls := 1
	if items[nextIndex].BlockedBy < 0 {
		if currentWeight+items[nextIndex].Weight <= allowedWeight {
			items[nextIndex].IsSelected = true
			test1Solution, test1Value, test1Calls = DoBranchAndBound(items, allowedWeight, nextIndex+1, bestValue,
				currentValue+items[nextIndex].Value,
				currentWeight+items[nextIndex].Weight,
				remainingValue-items[nextIndex].Value)
			if test1Value > bestValue {
				bestValue = test1Value
			}
		}
	}

	var test2Solution []Item
	var test2Value int
	var test2Calls int
	if currentValue+remainingValue-items[nextIndex].Value > bestValue {
		items[nextIndex].IsSelected = false
		test2Solution, test2Value, test2Calls = DoBranchAndBound(items, allowedWeight, nextIndex+1, bestValue,
			currentValue, currentWeight, remainingValue-items[nextIndex].Value)
		UnblockedItems(items[nextIndex], items)
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
