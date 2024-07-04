package data_structures

import (
	"fmt"
	"strings"
)

type BinaryTreeNode[T any] struct {
	Data  T
	Left  *BinaryTreeNode[T]
	Right *BinaryTreeNode[T]
}

func BuildTree() *BinaryTreeNode[string] {
	aNode := BinaryTreeNode[string]{
		Data: "A",
		Left: &BinaryTreeNode[string]{
			Data: "B",
			Left: &BinaryTreeNode[string]{
				Data: "D",
			},
			Right: &BinaryTreeNode[string]{
				Data: "E",
				Left: &BinaryTreeNode[string]{
					Data: "G",
				},
			},
		},
		Right: &BinaryTreeNode[string]{
			Data: "C",
			Right: &BinaryTreeNode[string]{
				Data: "F",
				Left: &BinaryTreeNode[string]{
					Data: "H",
					Left: &BinaryTreeNode[string]{
						Data: "I",
					},
					Right: &BinaryTreeNode[string]{
						Data: "J",
					},
				},
			},
		},
	}
	return &aNode
}

func (bt *BinaryTreeNode[T]) DisplayIndented(indent string, depth int) string {
	result := strings.Repeat(indent, depth) + fmt.Sprintf("%v\n", bt.Data)
	if bt.Left != nil {
		result += bt.Left.DisplayIndented(indent, depth+1)
	}
	if bt.Right != nil {
		result += bt.Right.DisplayIndented(indent, depth+1)
	}
	return result
}

func (bt *BinaryTreeNode[T]) Preorder() string {
	result := fmt.Sprintf("%v", bt.Data)
	if bt.Left != nil {
		result += " " + bt.Left.Preorder()
	}
	if bt.Right != nil {
		result += " " + bt.Right.Preorder()
	}
	return result
}

func (bt *BinaryTreeNode[T]) Inorder() string {
	result := ""
	if bt.Left != nil {
		result += bt.Left.Inorder() + " "
	}
	result += fmt.Sprintf("%v", bt.Data)
	if bt.Right != nil {
		result += " " + bt.Right.Inorder()
	}
	return result
}

func (bt *BinaryTreeNode[T]) PostOrder() string {
	result := ""
	if bt.Left != nil {
		result += bt.Left.PostOrder() + " "
	}
	if bt.Right != nil {
		result += bt.Right.PostOrder() + " "
	}
	result += fmt.Sprintf("%v", bt.Data)
	return result
}

func (bt *BinaryTreeNode[T]) BreadthFirst() string {
	result := ""
	return result
}
