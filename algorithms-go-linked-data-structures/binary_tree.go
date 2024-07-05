package data_structures

import (
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

type BinaryTreeNode[T constraints.Ordered] struct {
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

func (bt *BinaryTreeNode[T]) Insert(data T) {
	for {
		if data <= bt.Data {
			// add or move down the left branch
			if bt.Left == nil {
				bt.Left = &BinaryTreeNode[T]{
					Data: data,
				}
				break
			} else {
				bt.Left.Insert(data)
				break
			}
		} else {
			/// add or move down the right branch
			if bt.Right == nil {
				bt.Right = &BinaryTreeNode[T]{
					Data: data,
				}
				break
			} else {
				bt.Right.Insert(data)
				break
			}
		}
	}
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
	queue := CreateDoublyLinkedList[*BinaryTreeNode[T]](&BinaryTreeNode[T]{}, &BinaryTreeNode[T]{})
	queue.Enqueue(bt)
	for !queue.IsEmpty() {
		currentNode := queue.Dequeue()
		result += fmt.Sprintf("%v ", currentNode.Data.Data)
		if currentNode.Data.Left != nil {
			queue.Enqueue(currentNode.Data.Left)
		}
		if currentNode.Data.Right != nil {
			queue.Enqueue(currentNode.Data.Right)
		}
	}
	return result
}
