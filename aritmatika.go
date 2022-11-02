package main

import "fmt"

func main() {
	a := 20
	b := 10
	c := a + b
	d := c - a
	e := a * b
	f := e / a
	// g := e 

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	
	a %= 3
	fmt.Println(a)
	
	i := 21
	i++
	i++
	i++
	fmt.Println(i)
}