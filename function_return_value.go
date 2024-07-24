package main

import "fmt"

func getHello(name string) string {
	hello := "Hello" + name
	return hello
}
func main() {
	result := getHello("Egi")
	fmt.Println(result)

	fmt.Println("Budi")
	fmt.Println("joko")
}
