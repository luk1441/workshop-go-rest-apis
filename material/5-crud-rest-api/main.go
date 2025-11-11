package main

import (
	"fmt"
	"net/http"
)

func main() {
	const port = "8080"

	fmt.Println("Server listening on port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Server shut down")
	}
}
