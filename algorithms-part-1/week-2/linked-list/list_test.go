package list

import "testing"

func TestCreateList(t *testing.T) {
	list := CreateSingleLinkedList()
	if list == nil {
		t.Fatalf("failed to create list")
	}
}

func TestListSingleNode(t *testing.T) {
	list := CreateSingleLinkedList()
	if list == nil {
		t.Fatalf("failed to create list")
	}
	list.Append(32)
	node := list.Find(0)
	if node == nil {
		t.Fatalf("failed to find node in list at idx 0")
	}
	if node.Value != 32 {
		t.Fatalf("failed node value comparison")
	}
}

func TestListMultiNode(t *testing.T) {
	list := CreateSingleLinkedList()
	if list == nil {
		t.Fatalf("failed to create list")
	}
	list.Append(32)
	list.Append(64)
	list.Append(128)
	list.Prepend(16)
	list.Prepend(8)
	expectedOrder := []int{8, 16, 32, 64, 128}
	actualOrder := []int{}
	for node := list.Next(); node != nil; node = list.Next() {
		actualOrder = append(actualOrder, node.Value.(int))
	}
	if len(expectedOrder) != len(actualOrder) {
		t.Fatalf("expected order and actual order lengths do not match")
	}
	for idx, v := range expectedOrder {
		if actualOrder[idx] != v {
			t.Fatalf("idx: %d does not match expected value of %d", idx, v)
		}
	}
}
