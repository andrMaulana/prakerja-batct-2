package main

import "fmt"

//  buat struct untuk objek Mobil
type Car struct {
	Type   string
	FuelIn float64
}

// method untuk struct Mobil
func (c *Car) JarakTempuh() float64 {
	return c.FuelIn * 1.5 // jarak dalam mill
}

func main() {
	myCar := Car{
		Type:   "sedang",
		FuelIn: 21, // 30 liter bahan bakar
	}

	estimasiJarak := myCar.JarakTempuh() // memanggil method EstimateDistance()
	fmt.Printf("Jarak perkiraan yang bisa ditempuh : %.1f mill\n", estimasiJarak)
}
