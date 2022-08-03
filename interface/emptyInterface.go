/*

	=> golang bukan merupakan pemrograman berorientasi object (PBO)
	=> biasanya dalam PBO, ada satu data parent di puncak yang bisa dianggap sebagai semua implementasi data yang ada di bahasa pemrograman tersebut
	=> ex: di java ada java.lang.Object
	=> untuk menangani kasus seperti ini di golnag, kita dapat manggunakan interface kosong
	=> interface kosong merupakan interface yang tidak memiliki deklarasi method satupun, hal ini membuat secara otomatis semua tipe data akan menjadi semua implementasinya

*/

package main

import "fmt"

func Upps(number int) interface{} { // --> (number int, apapun interface{}) => parameter kedua dapat menerima tipe data apapun
	if number == 1 {
		return 1
	} else if number == 2 {
		return true
	} else {
		return "Upps"
	}
}

func main() {
	// var data interface{} = Upps(1)
	// var data interface{} = Upps(2)
	var data interface{} = Upps(3) // --> semua ipe data apapun dapat di terima dengan interface kosong

	fmt.Println(data)
}
