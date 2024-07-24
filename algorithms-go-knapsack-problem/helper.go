package knapsack

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	numItems  = 5
	minValue  = 1
	maxValue  = 10
	minWeight = 4
	maxWeight = 10
)

var allowedWeight int

type Item struct {
	Value      int
	Weight     int
	IsSelected bool
}

func MakeItems(numItems int, minValue int, maxValue int, minWeight int, maxWeight int) []Item {
	random := rand.New(rand.NewSource(1337))
	items := make([]Item, numItems)
	for i := 0; i < numItems; i++ {
		items[i] = Item{
			Value:      random.Intn(maxValue-minValue+1) + minValue,
			Weight:     random.Intn(maxWeight-minWeight+1) + minWeight,
			IsSelected: false,
		}
	}
	return items
}

func CopyItems(items []Item) []Item {
	newItems := make([]Item, len(items))
	copy(newItems, items)
	return newItems
}

func SumValues(items []Item, addAll bool) int {
	total := 0
	for i := 0; i < len(items); i++ {
		if addAll || items[i].IsSelected {
			total += items[i].Value
		}
	}
	return total
}

func SumWeights(items []Item, addAll bool) int {
	total := 0
	for i := 0; i < len(items); i++ {
		if addAll || items[i].IsSelected {
			total += items[i].Weight
		}
	}
	return total
}

func SolutionValue(items []Item, allowedWeight int) int {
	if SumWeights(items, false) > allowedWeight {
		return -1
	}
	return SumValues(items, false)
}

func PrintSelected(items []Item) {
	numPrinted := 0
	for i, item := range items {
		if item.IsSelected {
			fmt.Printf("%d(%d, %d)", i, item.Value, item.Weight)
		}
		numPrinted += 1
		if numPrinted > 100 {
			fmt.Println("...")
			return
		}
	}
	fmt.Println()
}

type Algorithm func([]Item, int) ([]Item, int, int)

func RunAlgorithm(alg Algorithm, items []Item, allowedWeight int) []Item {
	testItems := CopyItems(items)
	start := time.Now()
	solution, totalValue, functionCalls := alg(testItems, allowedWeight)
	elapsed := time.Since(start)
	fmt.Printf("Elapsed: %f\n", elapsed.Seconds())
	PrintSelected(solution)
	fmt.Printf("Value: %d, Weight: %d, Calls: %d\n\n", totalValue, SumWeights(solution, false), functionCalls)
	return solution
}
