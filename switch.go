package main

import "fmt"

func main() {
	name := "Safrizal"

	switch name {
	case "Saf":
		fmt.Println("Hello Saf")
	case "Sff":
		fmt.Println("Hi Sf")
	default:
		fmt.Println("Maaf")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama anda pas")
	}

	panjang := len(name)
	switch {
	case panjang > 10:
		fmt.Println("Nama anda terlalu panjang")
	case panjang > 5 :
		fmt.Println("Nama anda lumayan panjang")
	default:
		fmt.Println("Nama anda pas")
	}
}