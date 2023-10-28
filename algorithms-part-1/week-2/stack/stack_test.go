package stack

import "testing"

func TestCreateStack(t *testing.T) {
	stack := CreateStack()
	if stack == nil {
		t.Fatal("failed to create stack")
	}
}

func TestStackPush(t *testing.T) {
	stack := CreateStack()
	stack.Push(0)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	expectedStackList := []interface{}{5, 4, 3, 2, 1, 0}
	stackList := stack.List.ValuesHeadToTail()
	if len(expectedStackList) != len(stackList) {
		t.Fatal("expected stack list lengths do not match")
	}
	for idx, value := range expectedStackList {
		if stackList[idx] != value {
			t.Fatalf("expected stack list idx %d to contain value %d", idx, value)
		}
	}
}

func TestStackPop(t *testing.T) {
	stack := CreateStack()
	stack.Push(0)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	expectedStackList := []interface{}{5, 4, 3, 2, 1, 0}
	for _, value := range expectedStackList {
		stackValue := stack.Pop()
		if value != stackValue {
			t.Fatalf("expected value to be %d", stackValue)
		}
	}
}
