package main

import "fmt"

func main() {
	var panjang, lebar int

	fmt.Print("Masukkan panjang tabel perkalian: ")
	fmt.Scan(&panjang)
	fmt.Print("Masukkan lebar tabel perkalian: ")
	fmt.Scan(&lebar)

	fmt.Print("*\t")
	for i := 1; i <= panjang; i++ {
		fmt.Print(i, "\t")
	}
	fmt.Println()
	for i := 1; i <= lebar; i++ {
		fmt.Print(i, "\t")
		for j := 1; j <= panjang; j++ {
			fmt.Print(i*j, "\t")
		}
		fmt.Println()
	}
}
