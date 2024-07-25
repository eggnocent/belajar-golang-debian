package helper

import "fmt"

var version = "1.0.0"
var Application = "golang"

func sayGoodbye(name string) string {
	return "Good Bye" + name
}

func SayHello(name string) string {
	return "Hello" + name
}

func Contoh() {
	sayGoodbye("Egi")
	fmt.Println(version)
}
