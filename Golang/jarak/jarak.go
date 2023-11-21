package main

import "fmt"

func main() {
	var jarak, kecepatan, waktu float32

	fmt.Println("Program Menghitung Jarak Berdasarkan Kecepatan dan Waktu")
	fmt.Print("Masukkan kecepatan: ")
	fmt.Scan(&kecepatan)
	fmt.Print("Masukkan waktu: ")
	fmt.Scan(&waktu)

	jarak = kecepatan * waktu
	fmt.Println("Jika kecepatan", kecepatan, "km/jam dan waktu tempuh", waktu, "jam, maka jarak yang ditempuh adalah", jarak, "km.")
	// fmt.Printf("Jika kecepatan %v km/jam dan waktu tempuh %v jam, maka jarak yang ditempuh adalah %v km.\n", kecepatan, waktu, jarak)
	// fmt.Print("Jika kecepatan ", kecepatan, " km/jam dan waktu tempuh ", waktu, " jam, maka jarak yang ditempuh adalah ", jarak, " km.\n")
}
