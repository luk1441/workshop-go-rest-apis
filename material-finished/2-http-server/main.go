package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Welcome to Go Server!")
}

// exercise 2
func exampleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "This is the exampleHandler")
}

func main() {
	const port = "8080"

	http.HandleFunc("/", handler)

	// exercise 2
	http.HandleFunc("/example", exampleHandler)

	fmt.Println("Server listening on port " + port)
	http.ListenAndServe(":"+port, nil)
}
