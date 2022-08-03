/*
	=> struc merupakan sebuah template data yang digunakan untuk menggabungkakan 0 atau lebih tipe data dalam satu kesatuan (mirip object)
	=> Struct biasanya merupaka repreentasi data dapalm program aplikasi yang kita buat
	=> data sctruct disimpan dlaam field\
	=> sederhananya struct adalah kumpulan dari field

	=> ex : struct orang / kelas orang => ada nama, umur alamat dsb
*/

package main

import "fmt"

type Customer struct { // => biasanya orang golang menggunakan awalan huruf besar
	Name, Address string
	Age           int
}

func main() {

	// cara 1
	var agan Customer
	agan.Name = "gandi purna jen"
	agan.Address = "kota ambon"
	agan.Age = 26

	fmt.Println(agan)

	// cara 2
	mhilla := Customer{ // cara yang disarankan
		Name:    "karmila angriani tomia",
		Address: "kota ambon",
		Age:     28,
	}

	fmt.Println(mhilla)

	// cara 3
	kalam := Customer{"kalam salna jen", "kota sorong", 16} // posisiya harus sama, suatu saat struktur berubah maka error
	fmt.Println(kalam)
}
