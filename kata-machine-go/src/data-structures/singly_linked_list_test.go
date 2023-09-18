package data_structures

import "testing"

func TestSinglyLinkedList(t *testing.T) {
	l := SinglyLinkedList{}
	if l.Length != 0 {
		t.Fatalf("expected list length to be zero")
	}
	l.Append(1)
	if l.Length != 1 {
		t.Fatalf("expected list length to be one")
	}
	if l.Head != l.Tail {
		t.Fatalf("expected head and tail to be equal")
	}
	l.Append(2)
	if l.Length != 2 {
		t.Fatalf("expected length of list to be two")
	}
	if l.Head == l.Tail {
		t.Fatalf("expected list head and tail to be different")
	}
	// TODO(nick): add get test
	l.DumpInfo()
}
