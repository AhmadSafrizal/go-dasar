package main

import "fmt"

func main() {
	counter := 1

	for counter <= 99 {
		fmt.Print(" Ke", counter)
		counter++
	}

	for tambah := 1; tambah <=10; tambah++ {
		fmt.Print(" A", tambah)
	}

	slice := []string{"Ahmad", "Saf", "Rizal"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "Saf"
	person["job"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}