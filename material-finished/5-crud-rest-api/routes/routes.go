package routes

import (
	"net/http"

	"example.com/5-crud-rest-api/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/users", handlers.UsersHandler)
	http.HandleFunc("/user/", handlers.UserHandler)
}
