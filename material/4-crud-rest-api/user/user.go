package user

import (
	"errors"
	"fmt"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserInput struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{Id: 0, Name: "Max", Age: 30},
	{Id: 1, Name: "Anton", Age: 25},
}

func GetUsers() []User {
	return users
}

func GetUser(Id int) (User, error) {
	for _, u := range users {
		if u.Id == Id {
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User %d not found", Id)
}

func CreateUser(input UserInput) (User, error) {
	newID := len(users)
	newUser := User{Id: newID, Name: input.Name, Age: input.Age}
	if input.Name == "" {
		return User{}, errors.New("User couldnt be created")
	}
	users = append(users, newUser)
	return newUser, nil
}

func DeleteUser(id int) error {
	for i, u := range users {
		if u.Id == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User %d not found", id)
}

func UpdateUser(user User) (User, error) {
	for i, u := range users {
		if u.Id == user.Id {
			if user.Name == "" && user.Age <= 0 {
				return User{}, fmt.Errorf("User %d couldnt be updated", user.Id)
			}
			users[i].Name = user.Name
			users[i].Age = user.Age
			return users[i], nil
		}
	}
	return User{}, fmt.Errorf("User %d not found", user.Id)
}
