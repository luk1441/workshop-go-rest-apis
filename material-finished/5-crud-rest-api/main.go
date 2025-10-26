package main

import (
	"fmt"
	"net/http"

	"example.com/5-crud-rest-api/routes"
)

func main() {
	const port = "8080"
	routes.RegisterRoutes()

	fmt.Println("Server listening on port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Server shut down")
	}
}
