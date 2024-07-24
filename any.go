package main

import "fmt"

func ups() any {
	return "aaa"
}

func main() {
	var kosong any = ups()
	fmt.Println(kosong)
}
