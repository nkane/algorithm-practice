package sorting

import (
	"cmp"
	"time"
)

func QuickSort[T cmp.Ordered](list []T) []T {
	ShuffleSort[T](list, time.Now().UnixNano())
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

func QuickSortThreeWayPartition[T cmp.Ordered](list []T) []T {
	ShuffleSort[T](list, time.Now().Unix())
	QSortThreeWayPartition[T](list, 0, len(list)-1)
	return list
}

func QSortThreeWayPartition[T cmp.Ordered](list []T, low int, high int) {
	if high <= low {
		return
	}
	lt := low
	gt := high
	v := list[low]
	i := low
	for i <= gt {
		cmp := 0
		if list[i] < v {
			cmp = -1
		} else if list[i] > v {
			cmp = 1
		}
		switch cmp {
		case -1:
			list[lt], list[i] = list[i], list[lt]
			lt++
			i++
		case 1:
			list[i], list[gt] = list[gt], list[i]
			gt--
		default:
			i++
		}
	}
	QSortThreeWayPartition[T](list, low, lt-1)
	QSortThreeWayPartition[T](list, gt+1, high)
}
