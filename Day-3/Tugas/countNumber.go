package main

import (
	"fmt"
	"strconv"
)

func MunculSekali(input string) []int {
	uniqueNumbers := make(map[int]int)

	// menghitung jumlah kemunculan setiap angka pada input
	for _, char := range input {
		num, _ := strconv.Atoi(string(char))
		uniqueNumbers[num]++
	}

	// menyimpan angka yang hanya muncul satu kali
	var result []int
	for num, count := range uniqueNumbers {
		if count == 1 {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	input := "1234567778899" // contoh input
	result := MunculSekali(input)
	fmt.Println(result)
}
