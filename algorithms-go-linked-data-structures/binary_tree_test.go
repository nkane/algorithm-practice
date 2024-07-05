package data_structures

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestBinaryTree(t *testing.T) {
	tree := BuildTree()
	fmt.Println()
	fmt.Println(tree.DisplayIndented("  ", 0))

	assert.Assert(t, "A B D E G C F H I J" == tree.Preorder(), "pre-order: %+v\n", tree.Preorder())
	assert.Assert(t, "D B G E A C I H J F" == tree.Inorder(), "in-order: %+v\n", tree.Inorder())
	assert.Assert(t, "D G E B I J H F C A" == tree.PostOrder(), "post-order: %+v\n", tree.PostOrder())
	assert.Assert(t, "A B C D E F G H I J " == tree.BreadthFirst(), "breadth-first: %v\n", tree.BreadthFirst())
}
