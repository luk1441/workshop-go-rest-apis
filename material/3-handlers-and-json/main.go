package main

import (
	"fmt"
	"net/http"
)

func encodeHandler(w http.ResponseWriter, r *http.Request) {}

func decodeHandler(w http.ResponseWriter, r *http.Request) {}

func main() {
	const port = "8080"

	http.HandleFunc("/encode", encodeHandler)
	http.HandleFunc("/decode", decodeHandler)

	fmt.Println("Server listening on port " + port)
	http.ListenAndServe(":"+port, nil)
}
