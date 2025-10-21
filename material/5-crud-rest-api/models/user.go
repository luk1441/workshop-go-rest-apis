package models

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{Id: 1, Name: "Max"},
	{Id: 2, Name: "Anton"},
}

func GetUsers() []User {
	return users
}

func GetUser(id int) (User, bool) {
	for _, u := range users {
		if u.Id == id {
			return u, true
		}
	}
	return User{}, false
}

func CreateUser(name string) User {
	newID := len(users) + 1
	newUser := User{Id: newID, Name: name}
	users = append(users, newUser)
	return newUser
}

func DeleteUser(id int) bool {
	for i, u := range users {
		if u.Id == id {
			users = append(users[:i], users[i+1:]...)
			return true
		}
	}
	return false
}

func UpdateUser(id int, newName string) (User, bool) {
	for i, u := range users {
		if u.Id == id {
			users[i].Name = newName
			return users[i], true
		}
	}
	return User{}, false
}
