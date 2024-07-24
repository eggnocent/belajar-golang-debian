package main

import "fmt"

func main() {
	x := 5

	var p *int

	p = &x

	fmt.Println("Nilai x:", x)
	fmt.Println("Alamat x:", &x)

	fmt.Println("Nilai yang ditunjuk oleh p:", *p)
}
