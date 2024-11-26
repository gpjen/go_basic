package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

func NewNode(val int) *Node {
	return &Node{Value: val, Next: nil, Prev: nil}
}

func (n *Node) Append(val int) *Node {
	n.Next = NewNode(val)
	n.Next.Prev = n
	return n.Next
}

func (n *Node) Prepend(val int) *Node {
	n.Prev = NewNode(val)
	n.Prev.Next = n
	return n.Prev
}

func (n *Node) PopNext() *Node {
	n.Next = nil
	return n
}

func (n *Node) PopPrev() *Node {
	n.Prev = nil
	return n
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

func (ll *LinkedList) instantiate(val int) {
	newNode := NewNode(val)
	ll.Head = newNode
	ll.Tail = newNode
	ll.Length++
}

func (ll *LinkedList) reset() {
	ll.Head = nil
	ll.Tail = nil
	ll.Length = 0
}

func (ll *LinkedList) Append(val int) {
	if ll.Head == nil {
		ll.instantiate(val)
		return
	}

	ll.Tail = ll.Tail.Append(val)
	ll.Length++
}

func (ll *LinkedList) Prepend(val int) {
	if ll.Head == nil {
		ll.instantiate(val)
		return
	}

	ll.Head = ll.Head.Prepend(val)
	ll.Length++
}

func (ll *LinkedList) PopNext() {
	if ll.Head == nil {
		fmt.Println("Linkedlist is empty...")
		return
	}

	if ll.Head == ll.Tail {
		ll.reset()
		return
	}

	ll.Tail = ll.Tail.Prev.PopNext()
	ll.Length--
}

func (ll *LinkedList) PopPrev() {
	if ll.Head == nil {
		fmt.Println("Linkedlist is empty...")
		return
	}

	if ll.Head == ll.Tail {
		ll.reset()
		return
	}

	ll.Head = ll.Head.Next.PopPrev()
	ll.Length--
}

func (ll *LinkedList) Print() {
	root := ll.Head

	for root != nil {
		fmt.Printf(" -> %d", root.Value)
		root = root.Next
	}

	fmt.Println(" -> nil \nlength : ", ll.Length)
}

// func (ll *LinkedList)

func main() {
	fmt.Println("----- dobli linkedlist -----")

	myLL := NewLinkedList()
	myLL.Append(1)
	myLL.Append(2)
	myLL.Append(3)
	myLL.PopPrev()
	myLL.PopNext()
	myLL.Prepend(4)
	myLL.Append(5)
	myLL.Append(6)
	myLL.Print()
}
