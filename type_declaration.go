package main

import "fmt"

func main() {
	type noKtp string

	var ktpEgi noKtp = "121212122121"
	var contohKtp noKtp = "2222222222"

	fmt.Println(ktpEgi)
	fmt.Println(contohKtp)

}
