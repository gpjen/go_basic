package typelist

import "fmt"

type Snode struct {
	Data int
	Next *Snode
}

type SLinkedList struct {
	Head   *Snode
	Tail   *Snode
	Length int
}

func NewSLinkedList() *SLinkedList {
	return &SLinkedList{}
}

// Tambahkan node baru ke akhir linked list
func (l *SLinkedList) Append(value int) {
	newNode := &Snode{Data: value}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}

	l.Length++
}

// Cetak seluruh node dalam linked list
func (l *SLinkedList) Print() {
	root := l.Head

	for root != nil {
		fmt.Printf("%d -> ", root.Data)
		root = root.Next
	}

	fmt.Print("nil")
	fmt.Printf(" (%d)\n", l.Length)

}
