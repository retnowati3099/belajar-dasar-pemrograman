package main

import "fmt"

func main() {
	// inisialisasi slice
	var animals = []string{"elephant", "cow", "goat", "sheep"}
	fmt.Println(animals[0])

	// hubungan slice dengan array & operasi slice
	var newAnimals = animals[0:2] // semua elemen mulai indeks ke-0 hingga sebelum indeks ke-2
	fmt.Println(newAnimals)

	fmt.Println(animals[0:0]) // menghasilkan slice kosong karena tidak ada elemen sebelum indeks ke-0

	fmt.Println(animals[4:4]) // menghasilkan slice kosong karena tidak ada elemen yang dimulai dari indeks ke-4

	fmt.Println(animals[:]) // semua elemen

	fmt.Println(animals[2:]) // semua elemen mulai dari indeks ke-2

	fmt.Println(animals[:2]) //semua elemen hingga sebelum indeks ke-2

	// fmt.Println(animals[4:0]) error pada penulisan, nilai depan harus lebih kecil atau sama dengan belakang
}

/*
Slice adalah reference elemen arrray
slice bisa dibuat atau dihasilkan dari manipulasi sebuah array ataupun slice lainnya
perubahan data di tiap elemen slice akan berdampak pada slice lain yang memiliki alamat memori yang sama
*/
