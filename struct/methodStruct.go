/*

	=> struct sama dengan tipe data lain, dapat digunakan sebagai parameter di function
	=> namun dapat menambahkan method kedalam function, sehingga seakan2 stcruc memiliki function
	=> method == fuction

*/

package main

import "fmt"

// struct person
type Person struct {
	Name, Address string
	Age           int
}

//method person
func (person Person) sayHello(name string) { // "person" bisa menggunakan nama lain
	fmt.Println("hello", name, ", myname is", person.Name)
}

func main() {
	agan := Person{
		Name:    "gandi purna jen",
		Address: "indonesia",
		Age:     26,
	}

	agan.sayHello("rizwan")

	fmt.Println(agan)
}
