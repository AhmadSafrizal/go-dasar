package main

import "fmt"

func main() {
	sayHello("Ahmad", "Safrizal")
}

func sayHello(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}
