package data_structures

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestDoublyLinkedList(t *testing.T) {
	dl := CreateDoublyLinkedList("TOP_SENTIENL", "BOTTOM_SENTIENL")
	dl.AddRange([]string{"a", "b", "c", "d", "e", "f"})
	fmt.Println(dl.ToString(" | "))

	assert.Assert(t, dl.IsEmpty() == false)
	assert.Assert(t, dl.Length() == 6)

	dl.TopSentinel.Next.AddBefore(&Node[string]{
		Data: "z",
	})
	dl.BottomSentinel.Previous.AddBefore(&Node[string]{
		Data: "x",
	})
	fmt.Println(dl.ToString(" | "))
}

func TestDoublyLinkListExample(t *testing.T) {
	dl := CreateDoublyLinkedList("TOP_SENTIENL", "BOTTOM_SENTIENL")
	animals := []string{
		"Ant",
		"Bat",
		"Cat",
		"Dog",
		"Elk",
		"Fox",
	}
	dl.AddRange(animals)
	fmt.Println(dl.ToString(" | "))
}

func TestDoublyLinkListDelete(t *testing.T) {
	dl := CreateDoublyLinkedList("TOP_SENTIENL", "BOTTOM_SENTIENL")
	animals := []string{
		"Ant",
		"Bat",
		"Cat",
		"Dog",
		"Elk",
		"Fox",
	}
	dl.AddRange(animals)
	fmt.Println(dl.ToString(" | "))
	dl.Enqueue("test-value-to-be-deleted")
	fmt.Println(dl.ToString(" | "))
	n := dl.Dequeue()
	fmt.Println(dl.ToString(" | "))
	fmt.Printf("%+v\n", n.Data)
}

func TestDoublyLinkedListQueue(t *testing.T) {
	queue := CreateDoublyLinkedList("TOP_SENTIENL", "BOTTOM_SENTIENL")
	queue.Enqueue("Agate")
	queue.Enqueue("Beryl")
	v := queue.Dequeue()
	assert.Assert(t, v != nil)
	assert.Assert(t, v.Data == "Agate")
	queue.Enqueue("Citrine")
	v = queue.Dequeue()
	assert.Assert(t, v != nil)
	assert.Assert(t, v.Data == "Beryl")
	v = queue.Dequeue()
	assert.Assert(t, v != nil)
	assert.Assert(t, v.Data == "Citrine")
	queue.Enqueue("Diamond")
	queue.Enqueue("Emerald")
	assert.Assert(t, queue.IsEmpty() == false)
}

func TestDoublyLinkedListDeque(t *testing.T) {
	deque := CreateDoublyLinkedList("TOP_SENTIENL", "BOTTOM_SENTIENL")
	deque.PushTop("Ann")
	deque.PushTop("Ben")
	v := deque.PopBottom()
	assert.Assert(t, v != nil)
	assert.Assert(t, v.Data == "Ann")
	deque.PushBottom("F-Cat")
	v = deque.PopBottom()
	assert.Assert(t, v != nil)
	assert.Assert(t, v.Data == "F-Cat")
	v = deque.PopBottom()
	assert.Assert(t, v != nil)
	assert.Assert(t, v.Data == "Ben")
	deque.PushBottom("F-Dan")
	deque.PushTop("Eva")
	for !deque.IsEmpty() {
		v = deque.PopBottom()
		assert.Assert(t, v != nil)
	}
}
