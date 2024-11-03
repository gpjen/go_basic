package typelist

import (
	"fmt"
	"strings"
)

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

type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head:   nil,
		Tail:   nil,
		Length: 0,
	}
}

func (ll *LinkedList) Push(value string) {
	newNode := NewNode(value)

	if ll.Head == nil {
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		ll.Tail = ll.Tail.Append(value)
	}

	ll.Length++
}

func (ll *LinkedList) Pop() {
	root := ll.Head

	if root == nil {
		fmt.Println("LinkedList is Empty...")
		return
	}

	if root.Next != nil {
		ll.Head = nil
		ll.Tail = nil
		ll.Length = 0
		return
	}

	for root.Next.Next != nil {
		root = root.Next
	}

	ll.Tail = root.Remove()
	ll.Length--
}

func (ll *LinkedList) Print() {
	root := ll.Head

	for root != nil {
		fmt.Println("Value :", root)
		root = root.Next
	}

	fmt.Println("Length :", ll.Length)
}

func (ll *LinkedList) Find(value string) *Node {
	root := ll.Head

	for root.Next != nil {
		if strings.EqualFold(root.Value, value) {
			return root
		}
		root = root.Next
	}
	return &Node{}

}
