package main

import "fmt"

type node struct {
	value int
	next  *node
}

func newnode(val int) *node {
	return &node{value: val, next: nil}
}

func (n *node) append(val int) *node {
	n.next = newnode(val)
	return n.next
}

func (n *node) remove() *node {
	n.next = nil
	return n
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

func (ll *linkedList) push(val int) {
	newNode := newnode(val)
	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		ll.tail = ll.tail.append(val)
	}

	ll.length++

}

func (ll *linkedList) pop() {
	if ll.head == nil {
		fmt.Println("Linked List Kosong...")
		return
	}

	root := ll.head
	if root.next == nil {
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
		fmt.Println(root.value)
		root = root.next
	}
}

func (ll *linkedList) find(val int) *node {
	root := ll.head

	for root != nil {
		if root.value == val {
			return root
		}
		root = root.next
	}

	return nil
}

func main() {
	myLl := newLinkedList()
	myLl.push(21)
	myLl.push(31)
	myLl.push(41)
	myLl.push(51)
	myLl.push(61)
	myLl.push(71)
	myLl.push(81)
	myLl.push(91)
	myLl.pop()
	myLl.pop()
	myLl.pop()
	myLl.pop()
	myLl.pop()
	myLl.pop()
	myLl.pop()
	myLl.print()

	fmt.Println("Head :", myLl.head)
	fmt.Println("Tail :", myLl.tail)
	fmt.Println("Length :", myLl.length)

	fmt.Println("Find :", myLl.find(21))
	fmt.Println("Find :", myLl.find(81))

}
