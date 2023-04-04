package main

import (
	"fmt"
	"strings"
)

func main() {
	// Array
	angka := [4]int{1, 2, 3, 4}
	// fmt.Println(angka)
	// fmt.Println(len(angka))
	// array for
	for i, nilai := range angka {
		fmt.Println("Perulangan ke-", i)
		fmt.Printf("nilai index dari %d adalah %d\n", i, nilai)
	}
	fmt.Println(strings.Repeat("#", 25))

	// Slice
	nilai := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(nilai)

	// Map
	nilai2 := map[string]string{"nama1": "andri"}
	fmt.Println(nilai2)
}
