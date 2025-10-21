package main

import "fmt"

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{Id: 0, Name: "Max"},
	{Id: 1, Name: "Anton"},
}

func getUsers() []User {
	return users
}

func getUser(Id int) (User, bool) {
	for _, u := range users {
		if u.Id == Id {
			return u, true
		}
	}
	return User{}, false
}

func createUser(name string) User {
	newID := len(users)
	newUser := User{Id: newID, Name: name}
	users = append(users, newUser)
	return newUser
}

func main() {
	fmt.Println("All users:", getUsers())

	user, found := getUser(1)
	if found {
		fmt.Println("Found user:", user)
	} else {
		fmt.Println("User not found")
	}

	newUser := createUser("Lisa")
	fmt.Println("Created new user:", newUser)

	fmt.Println("Updated users:", getUsers())
}
