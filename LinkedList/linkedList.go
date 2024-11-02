package main

import (
	"fmt"
	"strings"
)

type node struct {
	value string
	next  *node
}

func newNode(value string) *node {
	return &node{value: value, next: nil}
}

func (n *node) append(value string) *node {
	n.next = newNode(value)
	return n.next
}

func (n *node) remove() *node {
	n.next = nil
	return n.next
}

type linkedList struct {
	head   *node
	tail   *node
	length int
}

func newLinkedList() *linkedList {
	return &linkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (ll *linkedList) push(value string) {
	newNode := newNode(value)

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		ll.tail = ll.tail.append(value)
	}

	ll.length++
}

func (ll *linkedList) pop() {
	root := ll.head

	if root == nil {
		fmt.Println("LinkedList is Empty...")
		return
	}

	if root.next != nil {
		ll.head = nil
		ll.tail = nil
		ll.length = 0
		return
	}

	for root.next.next != nil {
		root = root.next
	}

	ll.tail = root.remove()
	ll.length--
}

func (ll *linkedList) print() {
	root := ll.head

	for root != nil {
		fmt.Println("Value :", root.value)
		root = root.next
	}

	fmt.Println("Length :", ll.length)
}

func (ll *linkedList) find(value string) *node {
	root := ll.head

	for root.next != nil {
		if strings.EqualFold(root.value, value) {
			return root
		}
		root = root.next
	}
	return &node{}

}

func main() {
	myLl := newLinkedList()

	myLl.push("pada")
	myLl.push("hari")
	myLl.push("mingu")
	myLl.print()

	findNode := myLl.find("HARIs")
	println("value :", findNode.value)
}
