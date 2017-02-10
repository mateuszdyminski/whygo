package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", jsonHandler) // HL
	http.ListenAndServe(":8090", nil) // HL
}

func jsonHandler(w http.ResponseWriter, req *http.Request) {
	js, err := json.Marshal(map[string]string {"hello": "wroclaw!"}) // HL
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") // HL
	w.Write(js) // HL
}
