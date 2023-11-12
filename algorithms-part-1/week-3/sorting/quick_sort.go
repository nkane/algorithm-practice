package sorting

import "cmp"

func QuickSort[T cmp.Ordered](list []T) []T {
	QSort[T](list, 0, len(list)-1)
	return list
}

func QSort[T cmp.Ordered](list []T, low int, high int) {
	if high <= low {
		return
	}
	pivot := Partition[T](list, low, high)
	QSort[T](list, low, pivot-1)
	QSort[T](list, pivot+1, high)
}

func Partition[T cmp.Ordered](list []T, low int, high int) int {
	i := low
	j := high + 1
	v := list[low]
	for {
		i++
		for list[i] < v {
			i++
			if i == high {
				break
			}
		}
		j--
		for v < list[j] {
			j--
			if j == low {
				break
			}
		}
		if i >= j {
			break
		}
		list[i], list[j] = list[j], list[i]
	}
	list[low], list[j] = list[j], list[low]
	return j
}
