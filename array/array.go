package main

import "fmt"

func main() {
	var array1 [5]string

	array1[0] = "satu"
	array1[1] = "dua"
	array1[2] = "tiga"

	fmt.Println(array1)

	array1[0] = "sinaga"

	fmt.Println(array1)

	// =========================

	var  array2 = [4]string{
		"ayam",
		"bebek",
		"maung",
	}

	fmt.Println(array2)
	fmt.Println("panjang array2", len(array2)) // len digunakan untuk mengetahui panjang array "bukan banyak data array"

}