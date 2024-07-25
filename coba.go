package main

import "fmt"

func main() {
	var e int = 10
	var ptr *int = &e

	fmt.Println(e)
	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = 20

	fmt.Println(e)
	fmt.Println(*ptr)
}
