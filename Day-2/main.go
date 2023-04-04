package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	fmt.Println("")
	var fruits = []string{"Apple", "Grape", "Banana", "Melon"}
	var newFruits = fruits[0:2]
	fmt.Println(newFruits)
}
