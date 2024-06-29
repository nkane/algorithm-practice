package solutions

import (
	"fmt"
	"math/rand"
)

// Make an array containing pseudorandom numbers in [0, max).
func make_random_array(num_items, max int) []int {
	arr := make([]int, num_items)
	for i := 0; i < num_items; i++ {
		arr[i] = rand.Intn(max)
	}
	return arr
}

// Print at most num_items items.
func print_array(arr []int, num_items int) {
	if len(arr) <= num_items {
		fmt.Println(arr)
	} else {
		fmt.Println(arr[:num_items])
	}
}

// Verify that the array is sorted.
func check_sorted(arr []int) {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			fmt.Println("The array is NOT sorted!")
			return
		}
	}

	fmt.Println("The array is sorted")
}
