package main

import (
	"fmt"
	"linkedlist2/typelist"
)

func main() {
	fmt.Println("start...")

	myLL := typelist.NewSLinkedList()
	myLL.Append(1)
	myLL.Append(2)
	myLL.Append(3)
	myLL.Append(4)
	myLL.Append(5)

	myLL.Print()
}
