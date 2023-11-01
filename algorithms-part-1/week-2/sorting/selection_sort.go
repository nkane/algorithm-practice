package sort

import "cmp"

func SelectionSort[T cmp.Ordered](list []T) []T {
	for i := 0; i < len(list); i++ {
		minIdx := i
		for j := i + 1; j < len(list); j++ {
			if list[minIdx] > list[j] {
				minIdx = j
			}
		}
		tmp := list[i]
		list[i] = list[minIdx]
		list[minIdx] = tmp
	}
	return list
}
