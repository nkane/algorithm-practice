package solutions

import (
	"fmt"
	"math/rand"
	"time"
)

func TestQuickSort() {
	rand.Seed(time.Now().UnixNano())

	// Get the number of items and maximum item value.
	var num_items, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&num_items)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display the unsorted array.
	arr := make_random_array(num_items, max)
	print_array(arr, 40)
	fmt.Println()

	// Sort and display the result.
	quicksort(arr)
	print_array(arr, 40)

	// Verify that it's sorted.
	check_sorted(arr)
}

func quicksort(arr []int) {
	// See if we should stop the recursion.
	if len(arr) < 2 {
		return
	}

	// Partition.
	pivot_i := partition(arr)

	// Recursively sort the two halves.
	quicksort(arr[0:pivot_i])
	quicksort(arr[pivot_i+1:])
}

// Partition the slice.
func partition(arr []int) int {
	// Set the lower and upper indexes.
	lo := 0
	hi := len(arr) - 1

	// Use the last element as the pivot.
	pivot := arr[hi]

	// Temporary pivot index
	i := lo - 1

	for j := lo; j < hi; j++ {
		// See if arr[j] <= pivot.
		if arr[j] <= pivot {
			// Move the temporary pivot index forward
			i = i + 1

			// Swap arr[i] and arr[j].
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Drop the pivot in the between the two halves.
	i = i + 1
	arr[i], arr[hi] = arr[hi], arr[i]

	// Return the pivot's index.
	return i
}
