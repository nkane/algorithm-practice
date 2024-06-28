package algorithms

func Sum(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	result := arr[0]
	result += Sum(arr[1:])
	return result
}

func Count(arr []int) int {
	if len(arr) == 1 {
		return 1
	}
	count := 0
	count = 1 + Count(arr[1:])
	return count
}

func Max(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	head := arr[0]
	max := Max(arr[1:])
	if head > max {
		max = head
	}
	return max
}
