package typelist

import (
	"fmt"
)

type Node struct {
	Data int
	Prev *Node
	Next *Node
}

type DoublyLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

// Menambahkan node di akhir daftar
func (dll *DoublyLinkedList) AddNode(data int) {
	newNode := &Node{Data: data}
	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		dll.Tail.Next = newNode
		newNode.Prev = dll.Tail
		dll.Tail = newNode
	}
	dll.Length++
}

// Menghapus node berdasarkan nilai
func (dll *DoublyLinkedList) RemoveNode(data int) {
	current := dll.Head
	for current != nil {
		if current.Data == data {
			if current.Prev != nil {
				current.Prev.Next = current.Next
			} else {
				dll.Head = current.Next // menghapus head
			}
			if current.Next != nil {
				current.Next.Prev = current.Prev
			} else {
				dll.Tail = current.Prev // menghapus tail
			}
			dll.Length--
			return
		}
		current = current.Next
	}
}

// Mencari node berdasarkan nilai
func (dll *DoublyLinkedList) FindNode(data int) *Node {
	current := dll.Head
	for current != nil {
		if current.Data == data {
			return current
		}
		current = current.Next
	}
	return nil
}

// Menampilkan semua node dalam daftar
func (dll *DoublyLinkedList) Display() {
	current := dll.Head
	for current != nil {
		fmt.Print(current.Data, " ")
		current = current.Next
	}
	fmt.Println()
}

// Mendapatkan panjang daftar
func (dll *DoublyLinkedList) GetLength() int {
	return dll.Length
}
