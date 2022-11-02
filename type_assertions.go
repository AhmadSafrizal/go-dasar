package main

import (
	"fmt"
)

func random() interface {} {
	return "Saf"
}

func main() {
	var result interface{} = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	switch value := result.(type) {
	case int:
		fmt.Println("Value", value, "is int")
	case string:
		fmt.Println("Value", value, "is string")
	default:
		fmt.Println("Unknowntype")
	}
}