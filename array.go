package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Jono"
	names[1] = "Su"
	names[2] = "S"

	var values = [3]int{
		30,
		90,
		100,
	}

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	fmt.Println("")

	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println("")

	fmt.Println(names)

	fmt.Println("")

	fmt.Println(values)
}
