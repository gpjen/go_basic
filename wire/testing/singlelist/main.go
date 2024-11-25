package main

import (
	"fmt"
	"log"
)

type Node struct {
	Value int
	Next  *Node
}

func NewNode(val int) *Node {
	return &Node{Value: val, Next: nil}
}

func (n *Node) Append(val int) *Node {
	n.Next = NewNode(val)
	return n.Next
}

func (n *Node) Remove() *Node {
	n.Next = nil
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

func (ll *LinkedList) Push(val int) {
	if ll.Head == nil {
		newNode := NewNode(val)
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		ll.Tail = ll.Tail.Append(val)
	}

	ll.Length++

}

func (ll *LinkedList) Pop() {

	if ll.Head == nil {
		log.Println("Linkedlist Is Empty...")
		return
	}

	root := ll.Head

	if root.Next == nil {
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
		fmt.Printf("-> %d ", root.Value)
		root = root.Next
	}

	fmt.Println("-> nil, length :", ll.Length)

}

func main() {
	fmt.Printf("----- single linkedlist -----\n")

	myll := NewLinkedList()
	myll.Push(1)
	myll.Push(2)
	myll.Push(3)
	myll.Pop()
	myll.Push(4)
	myll.Push(5)
	myll.Print()

}
