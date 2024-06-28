package binary_search

import "golang.org/x/exp/constraints"

type SearchResult struct {
	Found      bool
	Iterations int
	Index      int
}

func BinarySearch[T constraints.Ordered](l []T, item T) *SearchResult {
	low := 0
	high := len(l) - 1
	found := false
	i := 0
	mid := 0
	for low <= high {
		i++
		mid = (low + high) / 2
		guess := l[mid]
		if guess == item {
			found = true
			break
		} else if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return &SearchResult{
		Found:      found,
		Iterations: i,
		Index:      mid,
	}
}
