package data_structures

type Node struct {
	Next  *Node       `json:"next"`
	Value interface{} `json:"value"`
}

type SinglyLinkedList struct {
	Head   *Node `json:"head"`
	Tail   *Node `json:"tail"`
	Length int   `json:"length"`
}
