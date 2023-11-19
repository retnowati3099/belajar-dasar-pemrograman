package main

import "fmt"

func main() {
	const PI float32 = 3.14 // deklarasi konstanta PI
	var (
		radius float32 //  deklarasi variabel radius
		luas   float32 // deklarasi variabel luas
	)
	println("Program Menghitung Luas Lingkaran")
	fmt.Println()
	fmt.Println("Masukkan radius: ")
	fmt.Scan(&radius)

	// proses menghitung luas lingkaran
	luas = PI * radius * radius

	// menampilkan hasil
	fmt.Printf("Lingkaran yang memiliki radius %v memiliki luas %v.\n", radius, luas)

}
