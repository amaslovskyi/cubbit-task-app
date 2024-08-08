package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Server is running on http://localhost:8888")
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
