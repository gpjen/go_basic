/*

	=> secara defeult di golang semua variabel itu diparsing by value, bukan by reference
	=> artinya, jika kita mengirim sebuah variable ke dalam function, method atau variabel lain, sebenarnya yang dikirimkan adalah duplikasi valuenya

	=> pointer adalah kemampuan membuat reference ke lokasi data di memori yang sama, tanpa menduplikasi data yag sudah ada
	=> sederhananya, dengan kemampuan ppointer, kita dapat membuat pass by reference

	=> kita dapat membuat pointer dengan operator &nama_variabel (buat berdasarkan variable yang telah ada)
	=> atau, function new (hanya mengembalikan pointer data kosong, artinya tidak ada nilai awal)
	   ex: alamat1 := new(Address)

*/

package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	addres1 := Address{"ambon", "maluku tengah", "indonesia"}
	addres2 := &addres1 // menggunakan tanda &nama_variabel
	addresBaru := new(Address)

	addres2.City = "masohi" //merubah adres2 maka adres1 ikut berubah
	// addres2 = &Address{"malang", "jawa timur", "indonesia"} // kalau di asign ulang addres2 maka addres1 tidak berubah
	// *addres2 = Address{"malang", "jawa timur", "indonesia"} // kalau ingin merubah addres2 maka harus menggunakan operator *
	// jadi siapapun yang mengacu ke variabel addres1 jika menggunakan * maka akan di pindahkan referensinya ke addres2

	fmt.Println("address 1 :", addres1)
	fmt.Println("address 2 :", addres2)
	fmt.Println("address baru :", addresBaru)
}
