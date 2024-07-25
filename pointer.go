package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	Address1 := Address{"Sleman", "D.I Yogyakarta", "Indonesia"}
	Address2 := Address1 //copy value

	Address2.City = "Klaten"
	fmt.Println(Address1)
	fmt.Println(Address2)
}
