package algorithms

func BinarySearch(key int, a []int) int {
	low := 0
	high := len(a) - 1
	for low <= high {
		mid := low + (high-low)/2
		if key < a[mid] {
			high = mid - 1
		} else if key > a[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
