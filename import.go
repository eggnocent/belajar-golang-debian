package main

import (
	"belajar-golang/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Egi")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version)
	// fmt.Println(helper.sayGoodbye("Egi"))

}