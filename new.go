package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var alamat1 *Address = new(Address)
	var alamat2 *Address = alamat1

	alamat2.Country = "Thailand"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
