/*

variadic function
=> parameter yang berada di posisi terakhir, memiliki kemampuan dijadikan sebuah varargs (variable argumen)
=> varargs artinya datanya dapat menerima lebih dari satu input, atau anggap sj semacam array
=> apa bedanya parameter biasa dengan tipe data array ?
		- parameter array, kita wajib membuat array terlebih dahulu sebelum mengirim ke function
		- jiga menggunakan varargs, kita dapat langsung mengirim datanya, jika lebih dari satu cukup menggunakan koma

*/

package main

import "fmt"

func main() {
	nilaiUjian := avgAll(10, 9, 5, 7, 10, 5, 8)

	fmt.Println("hasil ujian rata-rata : ", nilaiUjian)

	// jika sudah memiliki array / slide bisa gunakan cara ini
	tinggiBadanKelasB := []int{160, 155, 171, 166, 165, 164, 164}
	tinggiRataRata := avgAll(tinggiBadanKelasB...) // kirim slide dengan ...
	fmt.Println("tinggi kelas B rata-rata  : ", tinggiRataRata)

}

func avgAll(numbers ...int) int {
	total := 0
	for _, nilai := range numbers {
		total += nilai
	}
	return total / len(numbers)
}
