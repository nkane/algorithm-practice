package sorting

import (
	"cmp"
	"math/rand"
)

func ShuffleSort[T cmp.Ordered](list []T, seed int64) {
	rng := rand.New(rand.NewSource(seed))
	for i := len(list) - 1; i > 0; i-- {
		j := rng.Intn(i + 1)
		list[i], list[j] = list[j], list[i]
	}
}
