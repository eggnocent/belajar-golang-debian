package main

import "fmt"

func getCompliteName() (firstName string, middleName string, lastName string) {
	firstName = "Egi"
	middleName = "Wira"
	lastName = "Tama"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompliteName()
	fmt.Println(a, b, c)
}
