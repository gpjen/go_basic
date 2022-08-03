/*

=> di golang ada interface yang namanya error dan mempunyai func Error()
=> jadi bisa digunakan error bawaan golnag karna golang sudah menyedikan kontrak defeult.
=> sederhananya cukup buat data yang dimana fungctionnya mempunyai namanya Error()string

*/

package main

import (
	"errors" // import package errors
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) { // error adalah interface dapat di set nil
	if pembagi == 0 {
		return 0, errors.New("tidak dapat dibagi dengan NOL")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	var nilai, pembagi int
	fmt.Print("masukkan angka 1 :")
	fmt.Scan(&nilai)
	fmt.Print("\nmasukkan angka 2 :")
	fmt.Scan(&pembagi)

	hasil, err := Pembagian(nilai, pembagi)

	if err == nil {
		fmt.Println("Hasilnya :", hasil)
	} else {
		fmt.Println("peringatan!", err.Error())
	}
}
