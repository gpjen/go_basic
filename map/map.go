package main

import "fmt"
func main( ) {
	person := map[string]string {
		"firtName": "gandi",
		"midleName" : "purna",
		"lastName": "jen",
	}

	fmt.Println(person)
	fmt.Println("lastname : ", person["lastName"])

	//tambah data
	person["address"] = "kota ambon"
	fmt.Println("address : ", person["address"])

	/* function manipulasi map :
		len(map) 						=> untuk mendapatkan jumlah data di map
		make(map[TypeKey]TypeValue) 	=> membuat map baru
		delete(map, key)				=> menghapus data di map dengan  key
	*/

	book := make(map[string]string)

	book["title"] = "belajar golang"
	book["author"] = "gandi purna jen"
	book["year"] = "2026"
	book["upss"] = "salah"

	fmt.Println("book : ", book, "len : ", len(book))
	//hapus
	delete(book, "upss")
	fmt.Println("book : ", book, "len : ", len(book))


}


















