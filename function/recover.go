/*

=> recover merupakan function yang dapat kita gunkan untuk menangkap panic
=> Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan


=> simpan atau tangkap panic di defer function

*/

package main

import "fmt"

func main() {
	// runApp(true)
	runApp(false)
}

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("message error:", message)
	}
	fmt.Println("defer => APLIKASI SELESAI DIJALANKAN")
}

func runApp(Error bool) {
	defer endApp()
	if Error {
		panic("APLIKASI ERROR")
	}

	fmt.Println("APLIKASI RUNNING...")
}
