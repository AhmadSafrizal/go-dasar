package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
	// fmt.Println("Di sini ", man.Name)
}

func main() {
	saf := Man{"Saf"}
	saf.Married()
	fmt.Println(saf.Name)
}