package main

import (
	"fmt"
)

func main() {
	type NoKTP string
	type Married bool

	var noKtpSaya NoKTP = "3012892831"
	var marriedStatus Married = false
	fmt.Println(noKtpSaya)
	fmt.Println(marriedStatus)
}