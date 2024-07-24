package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil fucntion")
}

func runApplication() {
	defer logging()
	fmt.Println("Run App")
}

func main() {
	runApplication()
}
