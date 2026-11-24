package main

import "fmt"

func main() {
	var celsius, fahrenheit, kelvin float64

	fmt.Print("Masukkan suhu dalam Celsius: ")
	fmt.Scan(&celsius)

	fahrenheit = (celsius * 9 / 5) + 32
	kelvin = celsius + 273.15

	fmt.Println("Hasil Fahrenheit:", fahrenheit)
	fmt.Println("Hasil Kelvin:", kelvin)
}