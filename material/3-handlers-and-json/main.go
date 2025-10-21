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

func encodeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := Response{
		Message: "Hello",
		Status:  "Success",
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	const port = "8080"
	http.HandleFunc("/encode", encodeHandler)

	fmt.Printf("Server running at http://localhost:%v", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Server stopped: %v", err)
	}
}
