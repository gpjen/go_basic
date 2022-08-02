/*

Function As Parameter di golang,
artinya function dapat digunakan sebagai parameter untuk function lain

*/

package main

import "fmt"

func main() {

	var kata string
	fmt.Print("masukkan kata : ")
	fmt.Scan(&kata)
	sayHelloWithFillter(kata, spamFillter)
}

// type declaration
// jika fungsi terlalu panjang agak susah untuk menulisnya dalam parameter
// type declaration digukan sebagai alias sehingga memudahkan functions ebagai parameter
type Filter func(string) string

func sayHelloWithFillter(name string, fillter Filter) { // tanpa type parameter => (name string, filter func(string)string)
	fmt.Println("hallo, ", fillter(name))
}

func spamFillter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}
