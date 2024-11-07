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

// Remove node di akhir list
func (l *SLinkedList) Remove() {
	if l.Head == nil {
		fmt.Println("Linked list is empty...")
		return
	}

	if l.Head.Next == nil {
		l.Head = nil
		l.Tail = nil
		l.Length = 0
		return
	}

	curent := l.Head
	for curent.Next.Next != nil {
		curent = curent.Next
	}

	curent.Next = nil
	l.Tail = curent
	l.Length--
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
