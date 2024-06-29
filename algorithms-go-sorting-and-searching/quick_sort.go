package algorithms

import "golang.org/x/exp/constraints"

func QuickSort[T constraints.Ordered](array []T) []T {
	if len(array) < 2 {
		return array
	}
	pivot := array[0]
	less := []T{}
	greater := []T{}
	for i := 1; i < len(array); i++ {
		value := array[i]
		if value <= pivot {
			less = append(less, value)
		} else if value > pivot {
			greater = append(greater, value)
		}
	}
	less = QuickSort(less)
	greater = QuickSort(greater)
	result := append(less, pivot)
	result = append(result, greater...)
	return result
}
