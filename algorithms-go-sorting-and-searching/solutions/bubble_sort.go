package solutions

import (
	"fmt"
	"math/rand"
	"time"
)

func TestBubbleSort() {
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
	bubble_sort(arr)
	print_array(arr, 40)

	// Verify that it's sorted.
	check_sorted(arr)
}

// Use bubble sort to sort the array.
func bubble_sort(arr []int) {
	n := len(arr)
	swapped := true

	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			// If arr[i - 1] and arr[i] are out of order ...
			if arr[i-1] > arr[i] {
				// ... swap them.
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swapped = true
			}
		}
	}
}
