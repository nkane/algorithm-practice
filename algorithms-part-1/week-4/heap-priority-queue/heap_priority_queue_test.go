package heap_priority_queue

import (
	"testing"

	"gotest.tools/assert"
)

func TestCreateMaxPriorityQueue(t *testing.T) {
	pq := CreateMaxPriorityQueue[int](32)
	assert.Equal(t, pq.N, 0)
	assert.Equal(t, len(pq.Q), 33)
}

func TestMaxPriorityQueueInsertBasic(t *testing.T) {
	pq := CreateMaxPriorityQueue[rune](6)
	expectedOrder := []rune{0, 'Z', 'X', 'Y', 'A', 'C', 'B'}
	runes := []rune{'A', 'B', 'C', 'X', 'Y', 'Z'}
	for _, r := range runes {
		pq.Insert(r)
	}
	i := 0
	for _, r := range pq.Q {
		assert.Equal(t, r, expectedOrder[i])
		i++
	}
}

func TestMaxPriorityQueueInsertSwim(t *testing.T) {
	pq := CreateMaxPriorityQueue[rune](11)
	expectedOrder := []rune{0, 'T', 'S', 'R', 'N', 'P', 'O', 'A', 'E', 'I', 'G', 'H'}
	runes := []rune{'T', 'P', 'R', 'N', 'H', 'O', 'A', 'E', 'I', 'G'}
	for _, r := range runes {
		pq.Insert(r)
	}
	pq.Insert('S')
	i := 0
	for _, r := range pq.Q {
		assert.Equal(t, r, expectedOrder[i])
		i++
	}
}
