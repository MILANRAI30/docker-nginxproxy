package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{"status": "ok", "service": "1"}
	json.NewEncoder(w).Encode(resp)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{"message": "Hello from Service 1"}
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Service 1 is running on port 8001...")
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		log.Fatalf("Failed to start service: %v", err)
	}
}
