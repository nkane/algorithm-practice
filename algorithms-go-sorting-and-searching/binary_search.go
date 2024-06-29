package algorithms

import "golang.org/x/exp/constraints"

func BinarySearch[T constraints.Ordered](arr []T, low int, high int, x T) int {
	if high >= low {
		mid := low + (high-low)/2
		if arr[mid] == x {
			return mid
		}
		if arr[mid] > x {
			return BinarySearch(arr, low, mid-1, x)
		}
		return BinarySearch(arr, mid+1, high, x)
	}
	return -1
}
