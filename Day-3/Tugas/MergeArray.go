package main

import "fmt"

// Fungsi untuk menggabungkan dua array
func mergeArrays(arr1, arr2 []string) []string {
	merged := append(arr1, arr2...)
	return merged
}

// Fungsi untuk menghapus duplikasi nama
func removeDuplicates(arr []string) []string {
	unique := []string{}
	duplicate := map[string]bool{}

	for _, name := range arr {
		if !duplicate[name] {
			duplicate[name] = true
			unique = append(unique, name)
		}
	}

	return unique
}

func main() {
	// Array pertama
	array1 := []string{"Alice", "Bob", "Charlie", "David", "Eve"}

	// Array kedua
	array2 := []string{"Charlie", "David", "Frank", "Grace", "Hank"}

	// Mencetak array pertama
	fmt.Println("Array 1:", array1)

	// Mencetak array kedua
	fmt.Println("Array 2:", array2)

	// Menggabungkan kedua array
	mergedArray := mergeArrays(array1, array2)

	// Menghapus duplikasi nama
	uniqueArray := removeDuplicates(mergedArray)

	// Mencetak array hasil penggabungan dan penghapusan duplikasi
	fmt.Println("Hasil penggabungan dan penghapusan duplikasi:", uniqueArray)
}
