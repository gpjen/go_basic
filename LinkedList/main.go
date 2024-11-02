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

	println("---------------double-----------------")

	myLlDouble := typelist.NewLinkedlistDouble()

	myLlDouble.PushNext("di")
	myLlDouble.PushNext("suatu")
	myLlDouble.PushNext("hari")
	myLlDouble.PushNext("tanpa")
	myLlDouble.PushNext("sengaja")
	myLlDouble.PushNext("kita")

	// myLlDouble.PopPrev()
	myLlDouble.PopPrev()
	myLlDouble.PopPrev()
	myLlDouble.PopPrev()

	myLlDouble.PushPrev("bertemu")

	myLlDouble.Print()
}
