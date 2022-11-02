package main

import (
	"fmt"
)

func main() {

	// Keterangan const tidak bisa diubah
	const firstName string = "Ahmad"
	const lastName = "Safrizal"

	fmt.Println(firstName)
	fmt.Println(lastName)

	const (
		first string = "Satu"
		last = "Dua"
	)

	fmt.Println(first)
	fmt.Println(last)
}