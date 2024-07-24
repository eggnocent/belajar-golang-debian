package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "Egi"
	// person["address"] = "Sleman"

	person := map[string]string{
		"name":    "egi",
		"address": "sleman",
	}
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)
	// fmt.Println(person["posisi"]) salah, mencetak nilai default

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Egi"
	book["ups"] = "salah"

	fmt.Println(book)
	delete(book, "ups")
	fmt.Println(book)
}
