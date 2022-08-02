// break dan continue merupakan kata kunci yang dapat digunakan di perulangan
// break 		=> digunakan untuk menghentikan perulangan
// continue		=> mengentijna perulangan berjalna dan melanjutkan perulangan berikutnya
package main

import "fmt"

func main() {

	//break
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Println("perulangan ke : ", i)
	}

	//continue
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		fmt.Println("nilai ke : ", i)
	}

}
