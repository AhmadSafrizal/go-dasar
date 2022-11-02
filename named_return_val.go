package main

import "fmt"

func getCompleteName() (firstName, lastName string) {
	firstName = "Ahmad"
	lastName = "Safrizal"

	return
}

func main() {
	a, b := getCompleteName()
	fmt.Println(a)
	fmt.Println(b)
}