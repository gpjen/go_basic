/*

=> defer function adalah fungsi yang dapat kita jadwalkan untuk di eksekusi setelah sebuah function selesai di eksekusi
=> defer function akan selalu dieksekusi walaupun terjadi eror pada function yang di eksekusi

*/

package main

import "fmt"

func main() {

	// runApps(0) // tida dapat di bagi 0 menghasilkan error
	runApps(10)

}

func logging() {
	fmt.Println("selesai memanggil function")
}

func runApps(number int) {
	defer logging() // akan tetap dijalankan di akhir mesi runApps ada yang error
	fmt.Println("Running application...")
	for i := 1; i <= 10000; i++ {
		if i == 5000 {
			num := 10 / number
			fmt.Println(num)
		}
		fmt.Print(".")
	}
}
