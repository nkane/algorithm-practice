package data_structures

import "testing"

func TestSinglyLinkedListEmpty(t *testing.T) {
	l := SinglyLinkedList{}
	if l.Length != 0 {
		t.Fatalf("expected list length to be zero")
	}
	if l.Head != nil {
		t.Fatalf("expected list head to be nil")
	}
	if l.Tail != nil {
		t.Fatalf("expected list tail to be nil")
	}
	l.DumpInfo()
}

func TestSinglyLinkedListAppend(t *testing.T) {
	l := SinglyLinkedList{}
	if err := l.Append(1); err != nil {
		t.Fatalf("expected list to append value properly, error: %+v", err)
	}
	if l.Length != 1 {
		t.Fatalf("expected list length to be 1")
	}
	if l.Head != l.Tail {
		t.Fatalf("expected head and tail to be equal")
	}
	if err := l.Append(2); err != nil {
		t.Fatalf("expected list to append value properly, error: %+v", err)
	}
	if l.Length != 2 {
		t.Fatalf("expected length of list to be 2")
	}
	if l.Head == l.Tail {
		t.Fatalf("expected list head and tail to be different")
	}
	l.DumpInfo()
}

func TestSinglyLinkedListPrepend(t *testing.T) {
	l := SinglyLinkedList{}
	if err := l.Prepend(1); err != nil {
		t.Fatalf("expected list to prepend value properly, error: %+v", err)
	}
	if l.Length != 1 {
		t.Fatalf("expected list length to be 1")
	}
	if l.Head != l.Tail {
		t.Fatalf("expected head and tail to be equal")
	}
	if err := l.Prepend(2); err != nil {
		t.Fatalf("expected list to prepend value properly, error: %+v", err)
	}
	if l.Length != 2 {
		t.Fatalf("expected list length to be 2")
	}
	if l.Head.Value != 2 {
		t.Fatalf("expected head value to be 2")
	}
	if l.Tail.Value != 1 {
		t.Fatalf("expected head value to be 1")
	}
	if err := l.Prepend(1024); err != nil {
		t.Fatalf("expected list to prepend value properly, error: %+v", err)
	}
	if l.Length != 3 {
		t.Fatalf("expected list length to be 2")
	}
	if l.Head.Value != 1024 {
		t.Fatalf("expected head value to be 1024")
	}
	l.DumpInfo()
}

func TestSinglyLinkedListGet(t *testing.T) {
	l := SinglyLinkedList{}
	if err := l.Append(1); err != nil {
		t.Fatalf("expected list to append value properly, error: %+v", err)
	}
	if err := l.Append(2); err != nil {
		t.Fatalf("expected list to append value properly, error: %+v", err)
	}
	n0, err := l.Get(0)
	if err != nil {
		t.Fatalf("expected error to be nil for list get at index zero, error: %+v", err)
	}
	if n0.Value != 1 {
		t.Fatalf("expected first node value to be one got %+v instead", n0.Value)
	}
	n1, err := l.Get(1)
	if err != nil {
		t.Fatalf("expected error to be nil for list get at index one, error: %+v", err)
	}
	if n1.Value != 2 {
		t.Fatalf("expected second node value to be two got %+v instead", n1.Value)
	}
}
