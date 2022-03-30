package main

import (
	"fmt"
	"log"
	"net/http"
)

// Hello page handle
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Url path different
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	// Not a GET request case
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return

	}
	fmt.Fprintf(w, "Hello")
}

// Form page submit function
func formHandler(w http.ResponseWriter, r *http.Request) {
	// Parse data error
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err %v", err)
	}
	// Prints the form data to screen
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func main() {
	// File Server created (index file will be considered root)
	fileServer := http.FileServer(http.Dir("./static"))
	// Handle routes
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Server started at port 8080\n")
	// Start server in port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
