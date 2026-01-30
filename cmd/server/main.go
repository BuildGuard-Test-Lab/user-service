package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"status":"healthy","service":"user-service"}`)
	})

	http.HandleFunc("/api/v1/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"users":[]}`)
	})

	log.Println("Starting user-service on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
