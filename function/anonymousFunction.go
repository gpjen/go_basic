/*

 */

package main

import "fmt"

func main() {
	//anonymouse function 1
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("eko", blacklist)
	registerUser("anjing", blacklist)

	//anonymouse function 2
	registerUser("wardah", func(name string) bool {
		return name == "anjing" || name == "babi"
	})
	registerUser("babi", func(name string) bool {
		return name == "anjing" || name == "babi"
	})
}

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocket", name)
	} else {
		fmt.Println("Wellcome", name)
	}
}
