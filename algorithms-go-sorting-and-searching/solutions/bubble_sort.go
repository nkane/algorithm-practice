// Bubble sort

package main

import (
    "fmt"
    "math/rand"
    "time")

func main() {
    rand.Seed(time.Now().UnixNano())

    // Get the number of items and maximum item value.
    var num_items, max int;
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
        if arr[i - 1] > arr[i] {
            fmt.Println("The array is NOT sorted!")
            return
        }
    }

    fmt.Println("The array is sorted")
}

// Use bubble sort to sort the array.
func bubble_sort(arr []int) {
    n := len(arr)
    swapped := true

    for swapped {
        swapped = false
        for i := 1; i < n; i++ {
            // If arr[i - 1] and arr[i] are out of order ...
            if arr[i - 1] > arr[i] {
                // ... swap them.
                arr[i - 1], arr[i] = arr[i], arr[i - 1]
                swapped = true
            }
        }
    }
}
