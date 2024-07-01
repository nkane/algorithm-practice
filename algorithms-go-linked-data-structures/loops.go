package data_structures

func HasLoop[T any](list *LinkedList[T]) bool {
	fast := list.Sentinel.Next
	slow := list.Sentinel.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next
		if fast == slow || fast.Next == slow {
			return true
		}
		fast = fast.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
