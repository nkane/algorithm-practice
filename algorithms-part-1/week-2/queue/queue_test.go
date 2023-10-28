package queue

import "testing"

func TestCreateDeque(t *testing.T) {
	deque := CreateDeque()
	if deque == nil {
		t.Fatalf("failed to create deque")
	}
}

func TestDequeEnqueue(t *testing.T) {
	deque := CreateDeque()
	deque.Enqueue(0)
	deque.Enqueue(1)
	deque.Enqueue(2)
	deque.Enqueue(3)
	deque.Enqueue(4)
	deque.Enqueue(5)
	expectedDequeList := []interface{}{0, 1, 2, 3, 4, 5}
	dequeList := deque.List.ValuesHeadToTail()
	if len(expectedDequeList) != len(dequeList) {
		t.Fatalf("expected queue list lengths do not match")
	}
	for idx, value := range expectedDequeList {
		if dequeList[idx] != value {
			t.Fatalf("expected deque list idx %d to contain value %d", idx, value)
		}
	}
}

func TestDequeDequeue(t *testing.T) {
	deque := CreateDeque()
	deque.Enqueue(0)
	deque.Enqueue(1)
	deque.Enqueue(2)
	deque.Enqueue(3)
	deque.Enqueue(4)
	deque.Enqueue(5)
	expectedDequeList := []interface{}{0, 1, 2, 3, 4, 5}
	dequeList := deque.List.ValuesHeadToTail()
	if len(expectedDequeList) != len(dequeList) {
		t.Fatalf("expected queue list lengths do not match")
	}
	for _, value := range expectedDequeList {
		dequeValue := deque.Dequeue()
		if value != dequeValue {
			t.Fatalf("expected value to be %d", dequeValue)
		}
	}
}

func TestDequePeekFirst(t *testing.T) {
	deque := CreateDeque()
	deque.Enqueue(0)
	deque.Enqueue(1)
	deque.Enqueue(2)
	deque.Enqueue(3)
	deque.Enqueue(4)
	deque.Enqueue(5)
	expectedValue := 0
	dequePeekValue := deque.PeekFirst()
	if expectedValue != dequePeekValue {
		t.Fatal("expected value does not match peek value")
	}
}

func TestDequePeekLast(t *testing.T) {
	deque := CreateDeque()
	deque.Enqueue(0)
	deque.Enqueue(1)
	deque.Enqueue(2)
	deque.Enqueue(3)
	deque.Enqueue(4)
	deque.Enqueue(5)
	expectedValue := 5
	dequePeekValue := deque.PeekLast()
	if expectedValue != dequePeekValue {
		t.Fatal("expected value does not match peek value")
	}
}

func TestDequeAddFirst(t *testing.T) {
	deque := CreateDeque()
	deque.Enqueue(0)
	deque.Enqueue(1)
	deque.Enqueue(2)
	deque.Enqueue(3)
	deque.Enqueue(4)
	deque.Enqueue(5)
	deque.AddFirst(64)
	expectedValue := 64
	dequePeekValue := deque.PeekFirst()
	if expectedValue != dequePeekValue {
		t.Fatal("expected value does not match peek value")
	}
}

func TestDequeAddLast(t *testing.T) {
	deque := CreateDeque()
	deque.Enqueue(0)
	deque.Enqueue(1)
	deque.Enqueue(2)
	deque.Enqueue(3)
	deque.Enqueue(4)
	deque.Enqueue(5)
	deque.AddLast(64)
	expectedValue := 64
	dequePeekValue := deque.PeekLast()
	if expectedValue != dequePeekValue {
		t.Fatal("expected value does not match peek value")
	}
}

func TestDequeRemoveFirst(t *testing.T) {
	deque := CreateDeque()
	deque.Enqueue(128)
	deque.Enqueue(1)
	deque.Enqueue(2)
	deque.Enqueue(3)
	deque.Enqueue(4)
	deque.Enqueue(5)
	expectedValue := 128
	dequeValue := deque.RemoveFirst()
	if expectedValue != dequeValue {
		t.Fatal("expected value does not match peek value")
	}
	expectedDequeList := []interface{}{1, 2, 3, 4, 5}
	dequeList := deque.List.ValuesHeadToTail()
	if len(expectedDequeList) != len(dequeList) {
		t.Fatalf("expected queue list lengths do not match")

	}
	for idx, value := range expectedDequeList {
		if dequeList[idx] != value {
			t.Fatalf("expected deque list idx %d to contain value %d", idx, value)
		}
	}
}

func TestDequeRemoveLast(t *testing.T) {
	deque := CreateDeque()
	deque.Enqueue(0)
	deque.Enqueue(1)
	deque.Enqueue(2)
	deque.Enqueue(3)
	deque.Enqueue(4)
	deque.Enqueue(128)
	expectedValue := 128
	dequeValue := deque.RemoveLast()
	if expectedValue != dequeValue {
		t.Fatal("expected value does not match peek value")
	}
	expectedDequeList := []interface{}{0, 1, 2, 3, 4}
	dequeList := deque.List.ValuesHeadToTail()
	if len(expectedDequeList) != len(dequeList) {
		t.Fatalf("expected queue list lengths do not match")
	}
	for idx, value := range expectedDequeList {
		if dequeList[idx] != value {
			t.Fatalf("expected deque list idx %d to contain value %d", idx, value)
		}
	}
}
