/*

	=> walaupun method akan menempel di struct, tetapi sebenarnya data struct yang diakses methot merupakan pass by value
	=> sangat direkomendasikan menggunakan pointer dimethod, sehingga tidak boros memory karena harus selalu diduplikasi ketika memanggil method

*/

package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	agan := Man{
		Name: "gandi purna jen",
	}

	agan.Married()

	fmt.Println(agan)
}
