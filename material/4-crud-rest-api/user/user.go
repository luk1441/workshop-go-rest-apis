package user

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

func GetUser(Id int) User {
	for _, u := range users {
		if u.Id == Id {
			return u
		}
	}
	return User{}
}

func CreateUser(input UserInput) User {
	newID := len(users)
	newUser := User{Id: newID, Name: input.Name, Age: input.Age}
	if input.Name == "" {
		return User{}
	}
	users = append(users, newUser)
	return newUser
}

func DeleteUser(id int) {
	for i, u := range users {
		if u.Id == id {
			users = append(users[:i], users[i+1:]...)
			return
		}
	}
}

func UpdateUser(user User) User {
	for i, u := range users {
		if u.Id == user.Id {
			if user.Name == "" && user.Age <= 0 {
				return User{}
			}
			users[i].Name = user.Name
			users[i].Age = user.Age
			return users[i]
		}
	}
	return User{}
}
