package algorithms

import (
	"fmt"
	"math/rand"
	"time"
)

func MakeRandomSlice(numItems int, max int) []int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]int, numItems)
	for i := 0; i < len(result); i++ {
		result[i] = random.Intn(max)
	}
	return result
}

func PrintSlice(slice []int, numItems int) {
	if len(slice) <= numItems {
		fmt.Println(slice)
	} else {
		fmt.Println(slice[:numItems])
	}
}

func CheckSort(slice []int) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i-1] > slice[i] {
			fmt.Println("The slice is NOT sorted!")
			return false
		}
	}
	fmt.Println("The slice is sorted")
	return true
}
