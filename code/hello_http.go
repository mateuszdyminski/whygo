package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler) // HL
	http.ListenAndServe(":8080", nil) // HL
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello Wroclaw!") // HL
}
