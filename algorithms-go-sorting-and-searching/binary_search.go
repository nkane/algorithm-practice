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

func BinarySearchIterative[T constraints.Ordered](arr []T, x T) (int, int) {
	low := 0
	high := len(arr) - 1
	mid := 0
	numTest := 0
	for low <= high {
		numTest++
		mid = (low + high) / 2
		guess := arr[mid]
		if guess == x {
			return mid, numTest
		}
		if guess > x {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1, numTest
}
