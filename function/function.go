package main

import "fmt"

func main() {

	// pemanggilan function
	sayHello()
	sayHelloTo("gandi", "purna", 27)
	fmt.Println(getHello("ridwan kamil"))

	fmt.Println("========multiple return======")
	firstName, lastName, age := getFullName()
	fmt.Println(firstName, lastName, age)
	fmt.Println("========multiple namet return values======")
	a, b, c := myHobby()
	fmt.Println("hobi saya : ", a, b, c)
}

//function tanpa parameter
func sayHello() {
	fmt.Println("hello...")
}

//function dengan parameter
func sayHelloTo(firstName string, lastName string, age int) {
	fmt.Printf("hai, saya %s %s umur saya %d tahun \n", firstName, lastName, age)
}

//function return value
func getHello(name string) string {
	if name == "" {
		return "hai juga bro , salam kenal."
	}
	return "hai juga " + name + ", salam kenal."
}

//function retun multiple value
func getFullName() (string, string, int) {
	return "gandi", "purna jen", 26
}

//function namet return values
func myHobby() (hobi1 string, hobi2 string, hobi3 string) {
	hobi1 = "makan"
	hobi2 = "berenang"
	hobi3 = "ngoding"

	return
}
