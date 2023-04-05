package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe", "Qwe"}))
	// fmt.Println(Mapping([]string{}))

	// Slice yang berisi beberapa kata
	s := []string{"asd", "qwe", "asd", "adi", "qwe", "Qwe"}

	// Menghitung kemunculan setiap kata dalam slice
	wordCount := Mapping(s)

	// // Mencetak hasil
	fmt.Println("Kemunculan setiap kata dalam slice:")
	for word, count := range wordCount {
		fmt.Printf("kata %s muncul sebanyak %d kali\n", word, count)
	}
}

// Fungsi untuk menghitung kemunculan setiap kata dalam slice
// Mengembalikan map dengan key berupa kata dan value berupa jumlah kemunculan kata
func Mapping(s []string) map[string]int {
	wordCount := make(map[string]int)
	for _, word := range s {
		word = strings.ToLower(word) // Mengubah semua karakter menjadi lowercase
		wordCount[word]++
	}
	return wordCount // nilai yang di kembalikan berupa map
}
