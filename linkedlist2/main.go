package main

import (
	"fmt"
	"linkedlist2/typelist"
)

func main() {
	myLL := typelist.NewSLinkedList()
	myLL.Append(1)
	myLL.Append(2)
	myLL.Append(3)
	myLL.Remove()
	myLL.Remove()
	myLL.Append(4)
	myLL.Append(5)
	myLL.Append(6)
	myLL.Remove()
	myLL.Append(7)
	myLL.Print()

	node := myLL.Find(5)
	if node.Data > 0 {
		fmt.Println("find :", node.Data)
	}

	// ddl-------------------

	dll := &typelist.DoublyLinkedList{}

	dll.AddNode(10)
	dll.AddNode(20)
	dll.AddNode(30)

	fmt.Println("Daftar setelah penambahan:")
	dll.Display()

	dll.RemoveNode(20)
	fmt.Println("Daftar setelah penghapusan 20:")
	dll.Display()

	ddNode := dll.FindNode(30)
	if ddNode != nil {
		fmt.Println("Node ditemukan dengan data:", ddNode.Data)
	} else {
		fmt.Println("Node tidak ditemukan")
	}

	fmt.Println("Panjang daftar:", dll.GetLength())
}
