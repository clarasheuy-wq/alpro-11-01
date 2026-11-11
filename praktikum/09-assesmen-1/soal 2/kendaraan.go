package main

import "fmt"

func main() {
	var M, K, T int
	var kendaraan string
	M = 0
	K = 0
	T = 0

	for {
		fmt.Scan(&kendaraan)

		if kendaraan == "M" {
			M++
		} else if kendaraan == "K" {
			K++
		} else if kendaraan == "T" {
			T++
		} else {
			break
		}
	}

	fmt.Println("Tipe M: ", M)
	fmt.Println("Tipe K: ", K)
	fmt.Println("Tipe T: ", T)
}