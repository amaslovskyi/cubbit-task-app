package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Retrieve the API key from environment variables
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Println("API_KEY not set")
		return
	}

	// Register the handler function for the root path
	http.HandleFunc("/", helloHandler)

	// Start the HTTP server
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

// Handler function for incoming HTTP requests
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Respond with a greeting and the API key (Note: Exposing API keys is not recommended in production)
	fmt.Fprintf(w, "Hello, World! API Key: %s", os.Getenv("API_KEY"))
}
