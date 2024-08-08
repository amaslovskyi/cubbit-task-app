package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Println("Warning: API_KEY not set")
	}

	http.HandleFunc("/", helloHandler)

	fmt.Println("Server is running on http://localhost:8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! API Key: %s", os.Getenv("API_KEY"))
}
