package algorithms

import "golang.org/x/exp/constraints"

func QuickSortSlow[T constraints.Ordered](array []T) []T {
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
	less = QuickSortSlow(less)
	greater = QuickSortSlow(greater)
	result := append(less, pivot)
	result = append(result, greater...)
	return result
}

func QuickSort[T constraints.Ordered](a []T, lo int, hi int) {
	// ensure indices are in correct order
	if lo >= hi || lo < 0 {
		return
	}
	// partition array and get the pivot index
	p := Partition(a, lo, hi)
	// sort the two paritions
	QuickSort(a, lo, p-1) // left side of pivot
	QuickSort(a, p+1, hi) // right side of pivot
}

func Partition[T constraints.Ordered](a []T, lo int, hi int) int {
	pivot := a[hi] // choose the last element as pivot
	// temporary pivot index
	i := lo
	for j := lo; j <= hi-1; j++ {
		// if the current element is less than or equal to the pivot
		if a[j] <= pivot {
			// swap the current element with the element at the temporary pivot index
			a[i], a[j] = a[j], a[i]
			// move the temporary pivot index forward
			i++
		}
	}
	// swap the pivot with the last element
	a[i], a[hi] = a[hi], a[i]
	return i // the pivot index
}
