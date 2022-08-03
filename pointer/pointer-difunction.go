/*

	=> saat kita membuat parameter di function, secara default adalah pass by value, artinya data akan dicopy dan dikirimkan ke function.
	=> oleh kerna itu, saat kita merubah data di function maka data asli tidak berubah
	=> terkadang kita ingin merubah data asli tersebut
	=> untuk melakukan ini kita dapat menggunakan pointer di function
	=> untuk dijadikan parameter kita dapat menggunakan operator * di parameter

*/

package main

import "fmt"

type Address struct {
	City, Profince, Country string
}

func changeCountry(address *Address) {
	address.Country = "indonesia"
}

func main() {
	address1 := Address{
		City:     "ambon",
		Profince: "maluku tengah",
		Country:  "",
	}

	fmt.Println(address1)

	changeCountry(&address1) // merubah address1 by ponter memory

	fmt.Println(address1)
}
