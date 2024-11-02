package typelist

type Node struct {
	Value string
	Next  *Node
}

func NewNode(value string) *Node {
	return &Node{Value: value, Next: nil}
}

func (n *Node) Append(value string) *Node {
	n.Next = NewNode(value)
	return n.Next
}

func (n *Node) Remove() *Node {
	n.Next = nil
	return n.Next
}
