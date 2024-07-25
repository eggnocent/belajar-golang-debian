package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var Address1 Address = Address{"Sleman", "D.I Yogyakarta", "Indonesia"}
	var Address2 *Address = &Address1 // pointer
	Address2.City = "Klaten"
	fmt.Println(Address1) // ikut berubah
	fmt.Println(Address2) // berubah menjadi Klaten

	// Address2 = &Address{"Surakarta", "Jawa Tengah", "Indonesia"}
	*Address2 = Address{"Surakarta", "Jawa Tengah", "Indonesia"}
	fmt.Println(Address1)
	fmt.Println(Address2)
}
