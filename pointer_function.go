package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Uzbekistan"
}

func main() {
	var address *Address = &Address{}
	ChangeCountryToIndonesia(address)
	fmt.Println(address)
}
