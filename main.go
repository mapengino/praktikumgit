package main

import "fmt"

func main() {
	var pilihan int
	var celsius, hasil float64

	fmt.Println("=== Program Konversi Suhu ===")
	fmt.Println("1. Celsius ke Fahrenheit")
	fmt.Println("2. Celsius ke Kelvin")
	fmt.Println("3. Celsius ke Reamur")
	fmt.Print("Pilih menu (1-3): ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > 3 {
		fmt.Println("Pilihan tidak valid!")
		return
	}

	fmt.Print("Masukkan suhu dalam Celsius: ")
	fmt.Scan(&celsius)

	if celsius < -273.15 {
		fmt.Println("Suhu tidak valid! Minimal -273.15 Celsius.")
		return
	}

	if pilihan == 1 {
		hasil = (celsius * 9 / 5) + 32
		fmt.Println("Hasil:", hasil, "Fahrenheit")
	} else if pilihan == 2 {
		hasil = celsius + 273.15
		fmt.Println("Hasil:", hasil, "Kelvin")
	} else if pilihan == 3 {
		hasil = celsius * 4 / 5
		fmt.Println("Hasil:", hasil, "Reamur")
	}
}