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

	var drinks = []string{"freshwater", "milk", "tea", "coffee"}
	var aDrinks = drinks[0:3]
	var bDrinks = drinks[1:4]
	var aaDrinks = aDrinks[1:2]
	var baDrinks = bDrinks[0:1]

	fmt.Println(drinks)
	fmt.Println(aDrinks)
	fmt.Println(bDrinks)
	fmt.Println(aaDrinks)
	fmt.Println(baDrinks)

	// milk diubah menjadi juice
	baDrinks[0] = "juice"

	fmt.Println(drinks)
	fmt.Println(aDrinks)
	fmt.Println(bDrinks)
	fmt.Println(aaDrinks)
	fmt.Println(baDrinks)

	// fungsi len() untuk menghitung jumlah elemen slice
	fmt.Println(len(drinks))

	// fungsi append() untuk menambahkan elemen pada slice, diposisikan setelah indeks paling akhir

	var cDrinks = append(drinks, "wine")

	fmt.Println(drinks)
	fmt.Println(cDrinks)
}

/*
Slice adalah reference elemen arrray
slice bisa dibuat atau dihasilkan dari manipulasi sebuah array ataupun slice lainnya
perubahan data di tiap elemen slice akan berdampak pada slice lain yang memiliki alamat memori yang sama
*/
