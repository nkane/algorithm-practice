package data_structures

import (
	"testing"

	"gotest.tools/assert"
)

func TestSortedBinaryTree(t *testing.T) {
	bt := CreateSortedBinaryTree("C")
	bt.Insert("B")
	bt.Insert("A")
	bt.Insert("Z")

	find := bt.Find("Z")
	assert.Assert(t, find.Data == "Z")

	got := bt.Root.Inorder()
	assert.Assert(t, "A B C Z" == got, "got: %+v\n", got)
}

func TestSortedBinaryTreeBook(t *testing.T) {
	bt := CreateSortedBinaryTree("I")
	bt.Insert("G")
	bt.Insert("C")
	bt.Insert("E")
	bt.Insert("B")
	bt.Insert("K")
	bt.Insert("S")
	bt.Insert("Q")
	bt.Insert("M")
	bt.Insert("F")
	got := bt.Root.Inorder()
	assert.Assert(t, "B C E F G I K M Q S" == got, "got: %+v\n", got)
}
