package main

import "fmt"

// fungsi untuk pengecekan apakah bilangan yang di inputkan merupakan bilangan prima
func BilanganPrima(n int) bool { //fungsi ini menerima sebuah parameter dengan tipe data integer dan return value dengan tipe data boolean
	if n < 2 {
		return false // kondisi jika bernilai false
	}
	// melakukan iterasi bilangan yang akan di inputkan harus habis di bagi 2
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	// nilai yang akan di kembalikan
	return true
}

func main() {
	var n int
	fmt.Print("Input bilangan: ")
	fmt.Scan(&n)

	if BilanganPrima(n) {
		fmt.Printf("%d merupakan bilangan prima.\n", n)
	} else {
		fmt.Printf("%d merup bilangan prima.\n", n)
	}
}
