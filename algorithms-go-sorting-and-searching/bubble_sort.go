package algorithms

func BubbleSort(slice []int) {
	n := len(slice)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if slice[i] > slice[i+1] {
				slice[i+1], slice[i] = slice[i], slice[i+1]
				swapped = true
			}
		}
	}
}

/*
The bubble sort algorithm can be optimized by observing that the n-th pass finds
n-th largest element an dputs it into its final place. So, the inner loop can
avoid looking at the last n -1 items when running for the n-th time.
*/
func BubbleSortOptimized(slice []int) {
	n := len(slice)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if slice[i] > slice[i+1] {
				slice[i+1], slice[i] = slice[i], slice[i+1]
				swapped = true
			}
		}
		n = n - 1 // optimization
	}
}

func CocktailShakerSort(slice []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(slice)-1; i++ {
			if slice[i] > slice[i+1] { // test whether to two elements are in the wrong order
				slice[i+1], slice[i] = slice[i], slice[i+1]
				swapped = true
			}
		}
		if !swapped {
			break
		}
		swapped = false
		for i := len(slice) - 1; i > 0; i-- {
			if slice[i] > slice[i+1] { // test whether to two elements are in the wrong order
				slice[i+1], slice[i] = slice[i], slice[i+1]
				swapped = true
			}
		}
	}
}
