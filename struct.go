package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	var Egi Customer
	Egi.Name = "Egi Wiratama"
	Egi.Address = "Kenya"
	Egi.Age = 18

	fmt.Println(Egi.Name)
	fmt.Println(Egi.Address)
	fmt.Println(Egi.Age)

	//struct literal
	abin := Customer{
		Name:    "Abin",
		Address: "Mali",
		Age:     19,
	}
	fmt.Println(abin)

	//cara ke-2
	dito := Customer{"Dito", "Burundi", 19}
	fmt.Println(dito)

	dito.sayHello("Adit")
	abin.sayHello("Dino")
	Egi.sayHello("Fahri")
}
