package sort

import "cmp"

func ShellSort[T cmp.Ordered](list []T) []T {
	n := len(list)
	h := 1
	for h < (n / 3) {
		h = 3*h + 1
	}
	for h >= 1 {
		// insertion sort
		for i := h; i < n; i++ {
			for j := i; j >= h && list[j] < list[j-h]; j -= h {
				tmp := list[j]
				list[j] = list[j-h]
				list[j-h] = tmp
			}
		}
		h = h / 3
	}
	return list
}
