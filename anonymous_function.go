package main

import "fmt"

type Blacklist func(string) bool

func registeredUser(name string, Blacklist Blacklist) {
	if Blacklist(name) {
		fmt.Println("Youre are blocked", name)
	} else {
		fmt.Println("welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}
	registeredUser("Egi", blacklist)
	registeredUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
