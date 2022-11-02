package main

import "fmt"

func getGoodBye(name string) string {
	return "Gooog Bye " + name
}

func main() {
	bye := getGoodBye

	result := bye("Saf")

	fmt.Println(result)
	fmt.Println(bye("A"))
	fmt.Println(getGoodBye("B"))
}