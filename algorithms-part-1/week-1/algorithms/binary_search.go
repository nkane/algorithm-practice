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

// order growth: O(N^2)
func TwoSumSlow(a []int) int {
	count := 0
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if (a[i] + a[j]) == 0 {
				count++
			}
		}
	}
	return count
}

// order growth: O(N log N)
func TwoSumFast(a []int) int {
	// NOTE(nick):
	// - normally you would run a merge sort on the input array
	//  to allow for binary search to work, but for now we can
	//  assume the input array is sorted
	count := 0
	for i := 0; i < len(a); i++ {
		if BinarySearch(-a[i], a) > i {
			count++
		}
	}
	return count
}

// order growth: O(N^3)
func ThreeSumSlow(a []int) int {
	count := 0
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			for k := j + 1; k < len(a); k++ {
				if a[i]+a[j]+a[k] == 0 {
					count++
				}
			}
		}
	}
	return count
}

// order growth: O(N^2 log N)
func ThreeSumFast(a []int) int {
	// NOTE(nick):
	// - normally you would run a merge sort on the input array
	//  to allow for binary search to work, but for now we can
	//  assume the input array is sorted
	count := 0
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if BinarySearch(-a[i]-a[j], a) > j {
				count++
			}
		}
	}
	return count
}
