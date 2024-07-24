package main

import "fmt"

type User struct {
	Name, Email string
	Age         int
}

func (user User) GetName() string {
	return user.Name
}

func (user User) getEmail() string {
	return user.Email
}

func main() {
	var user User
	user.Name = "Egi"
	user.Email = "examp@pss.com"
	user.Age = 19

	fmt.Println(user.GetName())
	fmt.Println(user.getEmail())
}
