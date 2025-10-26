package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// exercise - 3
type RequestBody struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func encodeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := Response{
		Message: "Hello",
		Status:  "Success",
	}

	json.NewEncoder(w).Encode(response)
}

// exercise - 3
func decodeHandler(w http.ResponseWriter, r *http.Request) {
	var request RequestBody
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}
	fmt.Println(request)
}

func main() {
	const port = "8080"
	http.HandleFunc("/encode", encodeHandler)

	// exercise - 3
	http.HandleFunc("/decode", decodeHandler)

	fmt.Println("Server listening on port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Server shut down")
	}
}
