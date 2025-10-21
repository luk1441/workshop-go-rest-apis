package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"example.com/5-crud-rest-api/models"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := models.GetUsers()

	// exercise 5
	nameFilter := r.URL.Query().Get("name")
	var filteredUsers []models.User

	if nameFilter == "" {
		filteredUsers = users
	} else {
		for _, user := range users {
			lowerName := strings.ToLower(user.Name)
			lowerNameFilter := strings.ToLower(nameFilter)
			if strings.Contains(lowerName, lowerNameFilter) {
				filteredUsers = append(filteredUsers, user)
			}
		}
	}

	json.NewEncoder(w).Encode(filteredUsers)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/user/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
	}

	switch r.Method {
	case http.MethodGet:
		user, found := models.GetUser(id)
		if !found {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(user)

	case http.MethodPut:
		var reqBody struct {
			Name string `json:"name"`
		}
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		updatedUser, ok := models.UpdateUser(id, reqBody.Name)
		if !ok {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(updatedUser)

	case http.MethodDelete:
		ok := models.DeleteUser(id)
		if !ok {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"status": "deleted"})

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
