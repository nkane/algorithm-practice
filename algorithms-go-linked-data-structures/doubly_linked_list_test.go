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
