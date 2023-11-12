package sorting

import "cmp"

func MergeSort[T cmp.Ordered](list []T) []T {
	aux := make([]T, len(list))
	MSort(list, aux, 0, len(list)-1)
	return list
}

func MSort[T cmp.Ordered](list []T, aux []T, low int, high int) {
	if high <= low {
		return
	}
	mid := low + (high-low)/2
	MSort(list, aux, low, mid)
	MSort(list, aux, mid+1, high)
	Merge(list, aux, low, mid, high)
}

func Merge[T cmp.Ordered](list []T, aux []T, low int, mid int, high int) []T {
	// copy list
	for k := low; k <= high; k++ {
		aux[k] = list[k]
	}
	// merge lists
	i := low
	j := mid + 1
	for k := low; k <= high; k++ {
		if i > mid {
			list[k] = aux[j]
			j++
		} else if j > high {
			list[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			list[k] = aux[j]
			j++
		} else {
			list[k] = aux[i]
			i++
		}
	}
	return list
}

func IsSorted[T cmp.Ordered](list []T, low int, high int) bool {
	for i := low; i < high; i++ {
		if list[i] > list[i+1] {
			return false
		}
	}
	return true
}
