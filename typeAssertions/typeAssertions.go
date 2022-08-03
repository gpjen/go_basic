/*

   => typeAssertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
   => fitur ini sering sekali digunakan ketika kita bertemu dengan data interface kosong

*/

package main

import "fmt"

func Random() interface{} {
	return nil
}

func main() {
	var result interface{} = Random()
	// var resultString string = result.(string) //type assertions
	// fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("string :", value)
	case int:
		fmt.Println("number int :", value)
	case bool:
		fmt.Println("boolean :", value)
	default:
		fmt.Println("Unknown")
	}

}
