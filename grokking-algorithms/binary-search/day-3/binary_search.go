package binary_search

import "golang.org/x/exp/constraints"

func BinarySearch[T constraints.Ordered](l []T, x T) int {
	low := 0
	high := len(l) - 1
	mid := 0
	for low <= high {
		mid = (high + low) / 2
		guess := l[mid]
		if guess == x {
			return mid
		}
		if guess > x {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
