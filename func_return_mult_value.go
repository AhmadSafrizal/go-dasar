package main

import "fmt"

func getFullName() (string, string) {
	return "Ahmad", "Safrizal"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}