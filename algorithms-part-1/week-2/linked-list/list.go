package list

type Node struct {
	Next  *Node
	Value interface{}
}

type SingleLinkedList struct {
	Head *Node
	Tail *Node
}

func CreateSingleLinkedList(n *Node) *SingleLinkedList {
	list := SingleLinkedList{
		Head: n,
		Tail: n,
	}
	return &list
}

func (l *SingleLinkedList) Append(n *Node) {
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}
	l.Tail.Next = n
	l.Tail = n
}

func (l *SingleLinkedList) Prepend(n *Node) {
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}
	n.Next = l.Head
	l.Head = n
}

func (l *SingleLinkedList) Find(idx int) *Node {
	node := l.Head
	sIdx := 0
	for node != nil && sIdx < idx {
		node = node.Next
	}
	return node
}

func (l *SingleLinkedList) RemoveAt(idx int) {
	// TODO
}
