package main

import (
	"fmt"
	"net/http"

	"example.com/4-crud-rest-api/handlers"
)

func main() {
	const port = "8080"

	http.HandleFunc("/users", handlers.UsersHandler)
	http.HandleFunc("/user", handlers.CreateUserHandler)
	http.HandleFunc("/user/", handlers.UserHandler)

	fmt.Println("Server listening on port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Server shut down")
	}
}
