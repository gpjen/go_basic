package typelist

import (
	"fmt"
	"strings"
)

type MyValue struct {
	Name string
	Note string
}

type SinglyNode struct {
	Value *MyValue
	Next  *SinglyNode
}

func SinglyNewNode(value *MyValue) *SinglyNode {
	return &SinglyNode{Value: value, Next: nil}
}

func (n *SinglyNode) Append(value *MyValue) *SinglyNode {
	n.Next = SinglyNewNode(value)
	return n.Next
}

func (n *SinglyNode) Remove() *SinglyNode {
	n.Next = nil
	return n.Next
}

type SinglyLinkedlist struct {
	Head   *SinglyNode
	Tail   *SinglyNode
	Length int
}

func NewSinglyLinkedlist() *SinglyLinkedlist {
	return &SinglyLinkedlist{
		Head:   nil,
		Tail:   nil,
		Length: 0,
	}
}

func (ll *SinglyLinkedlist) Push(value *MyValue) {
	newSinglyNode := SinglyNewNode(value)

	if ll.Head == nil {
		ll.Head = newSinglyNode
		ll.Tail = newSinglyNode
	} else {
		ll.Tail = ll.Tail.Append(value)
	}

	ll.Length++
}

func (ll *SinglyLinkedlist) Pop() {
	if ll.Head == nil {
		fmt.Println("Linkedlist is empty...")
		return
	}

	if ll.Head.Next == nil {
		ll.Head = nil
		ll.Tail = nil
		ll.Length = 0
	}

	curent := ll.Head
	for curent.Next.Next != nil {
		curent = curent.Next
	}

	ll.Tail = curent.Remove()
	ll.Length--
}

func (ll *SinglyLinkedlist) Print() {
	root := ll.Head
	for root != nil {
		fmt.Printf("name : %s (%s)", root.Value.Name, root.Value.Note)
		root = root.Next
	}

	fmt.Println("length :", ll.Length)
}

func (ll *SinglyLinkedlist) Find(name string) *SinglyNode {
	root := ll.Head
	for root != nil {
		if strings.EqualFold(root.Value.Name, name) {
			return root
		}
		root = root.Next
	}

	return &SinglyNode{}
}
