package main

import (
	"fmt"
)

func main() {
	name := "Egi"

	switch name {
	case "Egi":
		fmt.Println("Halo Egi")
	case "Alfian":
		fmt.Println("Halo, Alfian")
	default:
		fmt.Println("Hi, boleh kenalan?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama sudah benar")
	}

	name = "karlesss"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("nama sudah benar")
	}
}
