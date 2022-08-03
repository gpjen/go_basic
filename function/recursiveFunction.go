/*
	function yang memanggil dirinya sendiri (recursive function)
*/

package main

import "fmt"

func main() {
	var number int
	fmt.Print("masukkan angka faktorial : ")
	fmt.Scan(&number)
	recursive := factorialRecursive(number)

	fmt.Printf("hasil faktorial %d adalah %d\n", number, recursive)
}

func factorialRecursive(number int) int {
	if number == 1 {
		//kondisi tidak memenuhi syarat memanggil diri sendiri
		return 1
	} else {
		// kondisi memanggil dirinya sendiri
		return number * factorialRecursive(number-1)
	}
}
