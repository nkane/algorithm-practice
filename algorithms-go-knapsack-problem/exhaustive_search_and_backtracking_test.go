package knapsack

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestExhaustiveSearch(t *testing.T) {
	items := MakeItems(5, minValue, maxValue, minWeight, maxWeight)
	allowedWeight := SumWeights(items, true)
	fmt.Println("*** Parameters ***")
	fmt.Printf("# items: %d\n", numItems)
	fmt.Printf("Total value: %d\n", SumValues(items, true))
	fmt.Printf("Total weight: %d\n", SumWeights(items, true))
	fmt.Printf("Allowed weight: %d\n", allowedWeight)
	fmt.Println()
	fmt.Println("*** Exhaustive Search ***")
	expected := []Item{
		{Value: 9, Weight: 5, IsSelected: true},
		{Value: 10, Weight: 4, IsSelected: true},
		{Value: 7, Weight: 5, IsSelected: true},
		{Value: 5, Weight: 6, IsSelected: true},
		{Value: 4, Weight: 5, IsSelected: true},
	}
	got := RunAlgorithm(ExhaustiveSearch, items, allowedWeight)
	assert.DeepEqual(t, expected, got)
}
