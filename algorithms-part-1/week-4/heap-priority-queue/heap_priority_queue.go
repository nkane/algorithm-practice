package heap_priority_queue

import "cmp"

type MaxPriorityQueue[T cmp.Ordered] struct {
	Q []T
	N int
}

func CreateMaxPriorityQueue[T cmp.Ordered](n int) *MaxPriorityQueue[T] {
	mq := MaxPriorityQueue[T]{
		N: n,
		Q: make([]T, n+1),
	}
	return &mq
}

func (q *MaxPriorityQueue[T]) Insert(value T) {
	q.N++
	if q.N > len(q.Q) {
		// TODO(nick): resize?
		q.Q = append(q.Q, value)
	} else {
		q.Q[q.N] = value
	}
}

func (q *MaxPriorityQueue[T]) Swim(k int) {
	for k > 1 && q.Less(k/2, k) {
		q.Exchange(k/2, k)
		k = k / 2
	}
}

func (q *MaxPriorityQueue[T]) Sink(k int) {
	for 2*k <= q.N {
		j := 2 * k
		if j < q.N && q.Less(j, j+1) {
			j++
		}
		if !q.Less(k, j) {
			break
		}
		q.Exchange(k, j)
		k = j
	}
}

func (q *MaxPriorityQueue[T]) DeleteMax() T {
	var zero T
	max := q.Q[1]
	q.N--
	q.Exchange(1, q.N)
	q.Q[q.N+1] = zero
	q.Sink(1)
	return max
}

func (q *MaxPriorityQueue[T]) Less(i int, j int) bool {
	if q.Q[i] > q.Q[j] {
		return false
	}
	return true
}

func (q *MaxPriorityQueue[T]) Exchange(i int, j int) {
	q.Q[i], q.Q[j] = q.Q[j], q.Q[i]
}
