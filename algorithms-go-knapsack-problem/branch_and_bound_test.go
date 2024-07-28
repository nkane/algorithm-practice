package knapsack

import (
	"fmt"
	"testing"
)

func TestBranchAndBound(t *testing.T) {
	items := MakeItems(5, minValue, maxValue, minWeight, maxWeight)
	allowedWeight := SumWeights(items, true)
	fmt.Println("*** Parameters ***")
	fmt.Printf("# items: %d\n", numItems)
	fmt.Printf("Total value: %d\n", SumValues(items, true))
	fmt.Printf("Total weight: %d\n", SumWeights(items, true))
	fmt.Printf("Allowed weight: %d\n", allowedWeight)
	fmt.Println()
	fmt.Println("*** Branch and Bound ***")
	got := RunAlgorithm(BranchAndBound, items, allowedWeight)
	fmt.Print(got)
}
