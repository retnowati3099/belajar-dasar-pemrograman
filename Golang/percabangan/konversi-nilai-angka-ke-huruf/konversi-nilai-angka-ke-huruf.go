// program konversi nilai angka yang diinput user menjadi huruf
package main

import "fmt"

func main() {
	var nilai int

	fmt.Print("Masukkan nilai kamu: ")
	fmt.Scan(&nilai)

	if nilai <= 100 && nilai >= 85 {
		fmt.Printf("Nilai kamu %d, kamu mendapat A.\n", nilai)
	} else if nilai >= 70 {
		fmt.Printf("Nilai kamu %d, kamu mendapat B.\n", nilai)
	} else if nilai >= 40 {
		fmt.Printf("Nilai kamu %d, kamu mendapat C.\n", nilai)
	} else if nilai >= 20 {
		fmt.Printf("Nilai kamu %d, kamu mendapat D.\n", nilai)
	} else {
		fmt.Printf("Nilai kamu %d, kamu mendapat E.\n", nilai)
	}
}
