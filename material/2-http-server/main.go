package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Go Server!")
}

func main() {
	const port = "8080"

	http.HandleFunc("/", handler)

	fmt.Println("Server listening on port " + port)
	http.ListenAndServe(":"+port, nil)
}
