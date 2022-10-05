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
	Engine   engine
}

type engine struct {
	Name string
	Hp int
}

func main() {
	users := user{FirstName: "John", LastName: "Doe", Email: "dev.dilshojdon@gmail.com", 
	Password: "123456", Age: 20,Engine: engine{Name: "V8",Hp: 500}}
	fmt.Println(users)
	fmt.Println(users.FirstName)
}