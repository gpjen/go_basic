/*

=> interface merupakan tipe data abstrak, tidak memiliki implementasi langsung
=> sebuah interfase berisikan definisi-definisi method
=> biasanya dinterface digunakann sebagai kontak

*/

package main

import "fmt"

// struct
type Person struct {
	Name, Address string
	Age           int
}

// methodstruct person
func (person Person) GetName() string {
	return person.Name
}

//interface
type HasName interface {
	GetName() string
}

func sayHello(hasname HasName) {
	fmt.Println("Hello,", hasname.GetName())
}

func main() {
	agan := Person{
		Name:    "gpjen",
		Address: "indonesia",
		Age:     26,
	}

	sayHello(agan)
}
