package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	egi := Man{"Egi"}
	egi.Married()

	fmt.Println(egi.Name)
}
