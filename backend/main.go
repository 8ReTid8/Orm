package main

import (
	"encoding/json"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello from Go backend",
	})
}

func main() {
	http.HandleFunc("/api/hello", helloHandler)

	println("Server running on :8080")

	http.ListenAndServe(":8080", nil)
}