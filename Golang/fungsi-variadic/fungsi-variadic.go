package main

import "fmt"

func main() {
	// pengisian parameter fungsi variadic menggunakan data slice
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var avg = calculate(numbers...)

	// atau

	// var avg = calculate(1, 2, 3, 4, 5, 6, 7, 8, 9)
	var msg = fmt.Sprintf("Rata - rata: %.2f", avg)
	fmt.Println(msg)
}

func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	var avg = float64(total) / float64(len(numbers))

	return avg
}

/*
Deklarasiparameter variadic sama dengan cara deklarasi variabel biasa, pembedana adalah pada variabel jenis ini ditambahkan tanda titik tiga (...) tepat setelah penulisan variabel (sebelum tipe data).
Semua nilai yang disisipkan sebagai parameter akan ditampung oleh variabel tersebut.

Pada contoh di atas,
numbers adalahsebuah variabel variadic.
Nilai tiap parameter bisa diakses seperti cara pengaksesan tiap elemen slice

Fungsi fmt.Sprintf() pada dasarnya sama dengan fmt.Printf(), hanya saja fungsi ini tidak menampilkan nilai, tetapi mengembalikan nilainya dalam bentuk string. Pada contoh di atas niali kembalian dari fmt.Sprintf() ditampung oleh variabel msg

Tipe data jika ditulis sebagai fungsi (ada tanda kurungnya) berguna untuk casting. Casting adalah teknik untuk konversi tipe sebauh data ke tipe lain.

OPerasi bilangan (perkalian, pembagian, dll) di GO hanya bisa dilakukan jika tipe datanya sejenis.

*/
