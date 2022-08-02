package main

import "fmt"

func main() {

	// test 1
	counter := 1
	for counter <= 5 {
		fmt.Println("loop ke : ", counter)
		counter++
	}

	//test 2

	slice := []string{"gandi", "purna", "jen"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// cara cepat mengambil semua data di array, slice map (for range)
	for _, name := range slice { // for index, as := range slice/map/array {} => index jika tidak digunakan makan pakai _
		// fmt.Printf("index ke %d = %s \n", index, name)
		fmt.Println("data : ", name)
	}

	// cara for rang map
	person := make(map[string]string)

	person["name"] = "karmila angriani tomia"
	person["address"] = "kota ambon"
	person["lastStudy"] = "D3 analist"

	for key, value := range person {
		fmt.Println(key, " = ", value)
	}
}
