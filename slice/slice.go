// slice merupakan potongan data dari array
// slice dan array selalu terkoneksi dimana slice adalah data yang mengakses sebagian atau seluruh array
// ex: array[low:hight],array[:hight],array[low:],array[:]

//funclion pada slice
//		len(slice) => mandapatkan panjang,
// 		cap(slice) => mendapatkan kapasitas,
// 		append(slice, data) => membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas penuh maka akan membuat array baru,
// 		make([]tipedata, length, capacity) => membuat slice baru,
// 		copy(destination, source) => menyalin slilce dari source ke destination
package main

import "fmt"

func main(){
	var month = [...]string {
		"januari",
		"febuari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var slice = month[4:7]

	fmt.Println(slice)
	fmt.Println("panjang : ", len(slice))
	fmt.Println("kapasitas : ", cap(slice))

	// rubah nilai slice -> mounth ikut berubah
	slice[0] = "mei-baru-nih"

	fmt.Println("month :", month)

	//appent saat slice/array penuh maka akan membbuat array baru
	var slice2 = month[10:]  		//month tidak berubah
	// var slice2 = month[9:11]		//month berubah

	fmt.Println("slice2 : ", slice2)
	var slice3 = append(slice2, "gandiJen")
	slice3[1] = "ups" 
	fmt.Println("slice3 : ", slice3)
	fmt.Println("month : ", month)
	fmt.Println("slice2 : ", slice2)


	//------------------------------------------------------------------
	fmt.Println("--------------cara aman buat slice------------------")

	newSlice := make([]string , 4, 5)

	newSlice[0] = "gandi"
	newSlice[1] = "purna"
	newSlice[2] = "jen"

	fmt.Println("newSlice : ", newSlice)

	//copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)

	fmt.Println("copySlice : ", copySlice)

}