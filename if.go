package main

import "fmt"

func main() {
	name := "uzannn"

	if name == "Egi" {
		fmt.Println("Halo Egi")
	} else if name == "adit" {
		fmt.Println("Halo adit")
	} else if name == "abin" {
		fmt.Println("abin")
	} else if name == "uzann" {
		fmt.Println("halo uzann")
	} else {
		fmt.Println("Hi, Boleh kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama terlalu Pendek")
	}
}
