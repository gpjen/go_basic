/*
	=> closure merupakan kemampuan sebuah function berinteraksi dengan data-data disekiatarnya dalam scope yang sama
	=> salah gunakan closure dapat merubah data yang tidak sengaja dirubah
*/

package main

import "fmt"

func main() {
	counter := 0

	increment := func() {
		// counter := 0  // => conter fungsi increment
		fmt.Println("increment ++")
		counter++ // jiga dalam fungsi incremen tidak ada counter maka akan merubah nilai counter di luar function
	}

	increment()
	increment()

	fmt.Println(counter)

}
