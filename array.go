package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Egi"
	names[1] = "Wira"
	names[2] = "tama"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var valuess = [3]int{
		90,
		80,
		85,
	}
	fmt.Println(valuess)

	var values = [...]int{
		90,
		80,
		85,
		111,
		112,
		113,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values))

}
