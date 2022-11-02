package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 70

	var lulusNilaiAkhir = nilaiAkhir >= 80
	var lulusAbsensi = absensi >= 80

	var lulus = lulusNilaiAkhir && lulusAbsensi
	// fmt.Println(lulus)

	if lulus == true {
		fmt.Println("Selamat anda lulus")
	} else {
		fmt.Println("Maaf anda belum lulus")
	}
}