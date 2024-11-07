package main

import (
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
}
