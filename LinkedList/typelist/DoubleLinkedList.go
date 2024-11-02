package typelist

import "fmt"

type NodeDouble struct {
	Value string
	Prev  *NodeDouble
	Next  *NodeDouble
}

func NewNodeDouble(value string) *NodeDouble {
	return &NodeDouble{
		Value: value,
		Prev:  nil,
		Next:  nil,
	}
}

func (n *NodeDouble) Append(value string) *NodeDouble {
	n.Next = NewNodeDouble(value)
	n.Next.Prev = n
	return n.Next
}

func (n *NodeDouble) Prepend(value string) *NodeDouble {
	n.Prev = NewNodeDouble(value)
	n.Prev.Next = n
	return n.Prev
}

func (n *NodeDouble) RemoveNext() *NodeDouble {
	n.Next = nil
	return n
}

func (n *NodeDouble) RemovePrev() *NodeDouble {
	n.Prev = nil
	return n
}

type LinkedListDouble struct {
	Head   *NodeDouble
	Tail   *NodeDouble
	Length uint
}

func NewLinkedlistDouble() *LinkedListDouble {
	return &LinkedListDouble{
		Head:   nil,
		Tail:   nil,
		Length: 0,
	}
}

func (ll *LinkedListDouble) Instatntiate(Value string) {
	n := NewNodeDouble(Value)
	ll.Head = n
	ll.Tail = n
	ll.Length++
}

func (ll *LinkedListDouble) Clear() {
	ll.Head = nil
	ll.Tail = nil
	ll.Length = 0
}

func (ll *LinkedListDouble) PushNext(value string) {
	if ll.Head == nil {
		ll.Instatntiate(value)
		return
	}

	ll.Tail = ll.Tail.Append(value)
	ll.Length++
}

func (ll *LinkedListDouble) PushPrev(value string) {
	if ll.Head == nil {
		ll.Instatntiate(value)
		return
	}

	ll.Head = ll.Head.Prepend(value)
	ll.Length++
}

func (ll *LinkedListDouble) PopNext() {
	if ll.Head == nil {
		return
	}

	if ll.Head == ll.Tail {
		ll.Clear()
		return
	}

	ll.Tail = ll.Tail.Prev.RemoveNext()
	ll.Length--
}

func (ll *LinkedListDouble) PopPrev() {
	if ll.Head == nil {
		return
	}

	if ll.Head == ll.Tail {
		ll.Clear()
		return
	}

	ll.Head = ll.Head.Next.RemovePrev()
	ll.Length--
}

func (ll *LinkedListDouble) Print() {
	if ll.Head == nil {
		fmt.Println("linkedlist Double id Empty...")
		return
	}

	root := ll.Head

	for root != nil {
		fmt.Println("Value :", root)
		root = root.Next
	}

	fmt.Println("lengt :", ll.Length)
}
