package algorithms

func LinearSearch[T comparable](l []T, target T) (int, int) {
	idx := -1
	numTest := 0
	for i := 0; i < len(l); i++ {
		numTest++
		if l[i] == target {
			idx = 1
			break
		}
	}
	return idx, numTest
}
