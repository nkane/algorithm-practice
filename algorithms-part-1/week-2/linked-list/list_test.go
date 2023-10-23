package list

import "testing"

func TestCreateList(t *testing.T) {
	list := CreateSingleLinkedList(nil)
	if list == nil {
		t.Fatalf("failed to create list")
	}
}

func TestListSingleNode(t *testing.T) {
	list := CreateSingleLinkedList(&Node{
		Next:  nil,
		Value: 32,
	})
	if list == nil {
		t.Fatalf("failed to create list")
	}
	node := list.Find(0)
	if node == nil {
		t.Fatalf("failed to find node in list at idx 0")
	}
	if node.Value != 32 {
		t.Fatalf("failed node value comparison")
	}
}
