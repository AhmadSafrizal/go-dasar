package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func (customer Customer) sayName(name string) {
	fmt.Println("Hello ", name, "My name is ", customer.Name)
}

func main() {
	var saf Customer

	saf.Name = "Safrizal"
	saf.Address = "SMG"
	saf.Age = 22

	saf.sayName("Jon")

	fmt.Println(saf.Name)
	fmt.Println(saf.Address)
	fmt.Println(saf.Age)

	joko := Customer {
		Name: "Joko",
		Address: "BGD",
		Age: 30,
	}

	fmt.Println(joko)

	budi := Customer{"Budi", "JKT", 22}
	fmt.Println(budi)
}