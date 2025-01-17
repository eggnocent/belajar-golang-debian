package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// Address1 := Address{"Sleman", "D.I Yogyakarta", "Indonesia"}
	// Address2 := Address1 //copy value

	// Address2.City = "Klaten"
	// fmt.Println(Address1) //tidak berubah
	// fmt.Println(Address2) //berubah menjadi Klaten

	var Address1 Address = Address{"Sleman", "D.I Yogyakarta", "Indonesia"}
	var Address2 *Address = &Address1 // pointer

	Address2.City = "Klaten"
	fmt.Println(Address1) // ikut berubah
	fmt.Println(Address2) // berubah menjadi Klaten
}
