package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Dockerized Go with Air! 12345")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server starting at port 8080...")
	http.ListenAndServe(":8080", nil)
}
