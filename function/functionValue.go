/*
di golang function dapat dibuat sebagai tipe data bahkan dapat disimpan sebagai sebuah variabel
*/

package main

import "fmt"

func main() {
	halo := sayHello // memasukkan fungsi ke sebuah variable
	//dapar dijaikan parameter function lain

	fmt.Println(halo("gandi purna jen"))
}

func sayHello(name string) string {
	return "hai, selamat malam " + name
}
