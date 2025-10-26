package user

import "fmt"

type User struct {
	Id   int
	Name string
}

func New(id int, name string) User {
	return User{Id: id, Name: name}
}

func (u User) Display() {
	fmt.Println(u.Name)
}

// exercise 1
func (u *User) ClearName() {
	u.Name = ""
}
