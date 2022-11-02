package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{"Grobogan", "Jawa Tengah", "Indonesia"}
	address2 := &address1
	address3 := &address1

	address2.City = "Semarang"

	*address2 = Address{"Surabaya", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	address4 := new(Address)
	address4.City = "Semarang"
	fmt.Println(address4)

	alamat := Address{
		City:     "Kediri",
		Province: "Jawa Timur",
		Country:  "",
	}
	changeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
}
