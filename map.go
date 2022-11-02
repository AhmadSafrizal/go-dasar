package main

import "fmt"

func main() {
	person := map[string]string {
		"name" : "Ahmad",
		"address": "Gubug",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	book := make(map[string]string)

	book["title"] = "Belajar golang"
	book["author"] = "Saf"
	book["channel"] = "SF"
	fmt.Println(book)
	fmt.Println(len(book))
	
	delete(book, "channel")
	fmt.Println(book)
	fmt.Println(len(book))
}