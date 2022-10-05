package main

import (
	"fmt"
)

type user struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
	Age       int
}

func main() {
	users := user{FirstName: "John", LastName: "Doe", Email: "dev.dilshojdon@gmail.com", Password: "123456", Age: 20}
	fmt.Println(users)
	fmt.Println(users.FirstName)
}