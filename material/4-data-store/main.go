package main

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{Id: 0, Name: "Max"},
	{Id: 1, Name: "Anton"},
}

func main() {

}
