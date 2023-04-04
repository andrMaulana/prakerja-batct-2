package main

import "fmt"

// Fungsi luasTrapesium menerima 3 parameter dg tipe data float64 dan mereturn float64
func luasTrapesium(sisiAtas, sisiBawah, tinggi float64) float64 {
	// melakukan perhitungan dari nilai yang di inputkan
	return (sisiAtas + sisiBawah) * tinggi / 2
}

func main() {
	// deklarasi variabel
	var sisiAtas, sisiBawah, tinggi float64
	// input nilai sisi atas
	fmt.Print("Masukkan panjang sisi atas: ")
	fmt.Scanln(&sisiAtas)

	// input nilai sisi bawah
	fmt.Print("Masukkan panjang sisi bawah: ")
	fmt.Scanln(&sisiBawah)

	// input nilai tinggi
	fmt.Print("Masukkan tinggi: ")
	fmt.Scanln(&tinggi)

	// menghitung luas Trapesium yang di tampung variable luas
	luas := luasTrapesium(sisiAtas, sisiBawah, tinggi)

	// mencetak luas trapesium
	fmt.Println("Luas trapesium adalah:", luas)
}
