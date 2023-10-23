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
	list := SingleLinkedList{}
	return &list
}

func (l *SingleLinkedList) Append(n *Node) {

}
