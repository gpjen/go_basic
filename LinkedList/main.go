package main

import "linkedlist/typelist"

func main() {
	myLl := typelist.NewLinkedList()

	myLl.Push("pada")
	myLl.Push("hari")
	myLl.Push("minggu")
	myLl.Push("rututut")
	myLl.Push("ayah")
	myLl.Print()

	findNode := myLl.Find("HARI")
	println("value :", findNode.Value)
}
