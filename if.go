package main

import "fmt"

func main() {
	name := "Sf"

	if name == "Saf" {
		fmt.Println("Hello Saf")
	} else if name == "Sf" {
		fmt.Println("Hi Sf")
	} else {
		fmt.Println("Maaf")
	}

	if length := len(name) ; length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama anda pas")	
	}
}