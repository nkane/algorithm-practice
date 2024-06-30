package solutions

import (
	"fmt"
	"math/rand"
	"time"
)

type Customer struct {
	id           string
	numPurchases int
}

func CountingSortTest() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display the unsorted slice.
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	sorted := countingSort(slice, max)
	printSlice(sorted, 40)

	// Verify that it's sorted.
	checkSorted(sorted)
}

// Make an slice containing pseudorandom
// customers with numPurchases in [0, max).
// The ith customer has id C<i> as in C103.
func makeRandomSlice(numItems, max int) []Customer {
	// Initialize a pseudorandom number generator.
	random := rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed

	slice := make([]Customer, numItems)
	for i := 0; i < numItems; i++ {
		id := fmt.Sprintf("C%d", i)
		numPurchases := random.Intn(max)
		slice[i] = Customer{id, numPurchases}
	}
	return slice
}

// Print at most numItems items.
func printSlice(slice []Customer, numItems int) {
	if len(slice) <= numItems {
		fmt.Println(slice)
	} else {
		fmt.Println(slice[:numItems])
	}
}

// Verify that the slice is sorted.
func checkSorted(slice []Customer) {
	for i := 1; i < len(slice); i++ {
		if slice[i-1].numPurchases > slice[i].numPurchases {
			fmt.Println("The slice is NOT sorted!")
			return
		}
	}

	fmt.Println("The slice is sorted")
}

// Sort an slice containing items in [0, max).
func countingSort(slice []Customer, max int) []Customer {
	numItems := len(slice)

	// Make an slice to hold counts.
	counts := make([]int, max)

	// Count the values.
	for i := 0; i < numItems; i++ {
		counts[slice[i].numPurchases]++
	}

	// Convert counts into counts of valus <=.
	for i := 1; i < max; i++ {
		counts[i] += counts[i-1]
	}

	// Count out the values.
	result := make([]Customer, numItems)
	for i := numItems - 1; i >= 0; i-- {
		// Move item i into position.
		num := slice[i].numPurchases
		result[counts[num]-1] = slice[i]
		counts[num] -= 1
	}

	return result
}
