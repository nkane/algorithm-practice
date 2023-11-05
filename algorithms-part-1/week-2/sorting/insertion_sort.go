package sort

import "cmp"

func InsertionSort[T cmp.Ordered](list []T) []T {
	for i := 0; i < len(list); i++ {
		for j := i; j > 0; j-- {
			if list[j] < list[j-1] {
				tmp := list[j]
				list[j] = list[j-1]
				list[j-1] = tmp
			}
		}
	}
	return list
}
