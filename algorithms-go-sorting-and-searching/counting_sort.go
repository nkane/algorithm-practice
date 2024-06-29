package algorithms

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Customer struct {
	ID           string
	NumPurchases int
}

func MakeCustomerSlice() []Customer {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]Customer, r.Int())
	for i := 0; i < len(result); i++ {
		id := fmt.Sprintf("C%d", r.Int())
		result[i] = Customer{
			ID:           id,
			NumPurchases: r.Int(),
		}
	}
	PrintSlice(result, len(result))
	return nil
}

func CheckCustomersSliceSorted(slice []Customer) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i-1].NumPurchases > slice[i].NumPurchases {
			fmt.Println("The slice is NOT sorted!")
			return false
		}
	}
	fmt.Println("The slice is sorted")
	return true
}

func CountingSort(a []Customer, digit int, radix int) []Customer {
	// a is a list to be sorted, radix is the base of the number system, digit
	// is the digit we want to sort by

	// create a list b which will be the sorted list
	b := make([]Customer, len(a))
	c := make([]int, len(a))

	// count the number of occurences of each digit in a
	for i := 0; i < len(a); i++ {
		digitOfAi := a[i].NumPurchases / int(math.Pow(float64(radix), float64(digit))) % radix
		c[digitOfAi] = c[digitOfAi] + 1
		// now c[i] is the value of the number of elements in a equal to i
	}

	// this for loop changes c to show the cumlative number of digits up to that index of c
	for j := 1; j < radix; j++ {
		// here c is modified to have the number of lements <= i
		c[j] = c[j] + c[j-1]
	}

	// to count down (go through a backwards)
	for m := len(a) - 1; m > 0; m-- {
		digitOfAi := a[m].NumPurchases / int(math.Pow(float64(radix), float64(digit))) % radix
		c[digitOfAi] = c[digitOfAi] - 1
		b[c[digitOfAi]] = a[m]
	}

	return b
}
