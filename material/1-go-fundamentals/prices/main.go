package main

import (
	"fmt"

	"example.com/prices/calculator"
	"example.com/prices/user"
)

func main() {
	sum := calculator.Add(2, 3)

	newUser := user.New(1, "Max")
	newUser.Display()

	fmt.Println(sum)
}
