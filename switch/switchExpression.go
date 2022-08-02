// biasanya digunakan untuk pengecekan pada satu variabel

package main

import "fmt"

func main() {
	var name string
	fmt.Print("masukkan nama : ")
	fmt.Scan(&name)

	switch name {
	case "agan":
		fmt.Println("hello agan")
	case "mhilla":
		fmt.Println("hello mhill")
	default:
		fmt.Println("hello World")
	}

	//short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama anda > 5 karakter")
	case false:
		fmt.Println("nama anda kurang panjang")
	}

	//switch tanpa kondisi
	var age int
	fmt.Print("masukkan umur : ")
	fmt.Scan(&age)

	switch {
	case age < 10:
		fmt.Println("bocil kematian")
	case age < 20:
		fmt.Println("anak baru gede")
	case age < 30:
		fmt.Println("remaja")
	default:
		fmt.Println("tua luu")
	}
}
