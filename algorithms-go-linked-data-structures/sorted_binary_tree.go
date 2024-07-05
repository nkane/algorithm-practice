package data_structures

import "golang.org/x/exp/constraints"

type SortedBinaryTree[T constraints.Ordered] struct {
	Root *BinaryTreeNode[T]
}

func CreateSortedBinaryTree[T constraints.Ordered](data T) *SortedBinaryTree[T] {
	return &SortedBinaryTree[T]{
		Root: &BinaryTreeNode[T]{
			Data: data,
		},
	}
}

func (bt *SortedBinaryTree[T]) Insert(data T) {
	bt.Root.Insert(data)
}

func (bt *SortedBinaryTree[T]) Find(target T) *BinaryTreeNode[T] {
	node := bt.Root
	for node != nil {
		if target == node.Data {
			return node
		}
		if target >= node.Data {
			// move right
			node = node.Right
		} else {
			// move left
			node = node.Left
		}
	}
	return node
}
