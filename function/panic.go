/*

=> panic function merupakan fungsi yang dapat kita panggil untuk menghentikan program
=> biasanya panic dipanggil ketika terjadi error saat program berjalan
=> saat panic function di panggil program akan terhenti namun defer akan tetap dieksekusi

*/

package main

import "fmt"

func main() {
	// runApp(true)  // => mengeksekusi blok panic("APLIKASI ERROR")
	runApp(false)
}

func endApp() {
	fmt.Println("defer enapps => aplikasi selesai di jalankan")
}

func runApp(Error bool) {
	defer endApp()
	if Error {
		panic("APLIKASI ERROR")
	}

	fmt.Println("APLIKASI BERJALAN")
}
