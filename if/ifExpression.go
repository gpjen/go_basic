package main

import "fmt"

func main() {
	var name string 
	
	fmt.Print("input your name : ")
	fmt.Scan(&name)

	if name == "gandijen" {
		fmt.Println("hello gandi")
	} else if name == "mhilla" {
		fmt.Println("hello mhill")
	} else {
		fmt.Println("hello world")
	}


	// short statement if

	if length := len(name); length < 2 {
		fmt.Println("nama apa ini kurang dari 3")
	} else if length < 5 {
		fmt.Println("nama ini hampir mantap  < 5")
	} else {
		fmt.Println("nama apa ini panjang beut")
	}
}