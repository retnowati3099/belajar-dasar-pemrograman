package main

import (
	"fmt"
	"strings"
)

func main() {
	// pengisian parameter fungsi variadic menggunakan data slice
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var avg = calculate(numbers...)

	// atau

	// var avg = calculate(1, 2, 3, 4, 5, 6, 7, 8, 9)
	var msg = fmt.Sprintf("Rata - rata: %.2f", avg)
	fmt.Println(msg)

	// fungsi dengan parameter biasa & variadic
	myHobby("Retno Wati", "coding", "reading")

	var hobbies = []string{"coding", "reading"}
	myHobby("Retno Wati", hobbies...)
}

func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	var avg = float64(total) / float64(len(numbers))

	return avg
}

func myHobby(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")
	fmt.Printf("Hello, my name is %s\n", name)
	fmt.Printf("My hobbies are %s\n", hobbiesAsString)
}

/*
Deklarasiparameter variadic sama dengan cara deklarasi variabel biasa, pembedana adalah pada variabel jenis ini ditambahkan tanda titik tiga (...) tepat setelah penulisan variabel (sebelum tipe data).
Semua nilai yang disisipkan sebagai parameter akan ditampung oleh variabel tersebut.

Pada contoh di atas,
numbers adalahsebuah variabel variadic.
Nilai tiap parameter bisa diakses seperti cara pengaksesan tiap elemen slice

Fungsi fmt.Sprintf() pada dasarnya sama dengan fmt.Printf(), hanya saja fungsi ini tidak menampilkan nilai, tetapi mengembalikan nilainya dalam bentuk string. Pada contoh di atas niali kembalian dari fmt.Sprintf() ditampung oleh variabel msg

Tipe data jika ditulis sebagai fungsi (ada tanda kurungnya) berguna untuk casting. Casting adalah teknik untuk konversi tipe sebauh data ke tipe lain.

Operasi bilangan (perkalian, pembagian, dll) di GO hanya bisa dilakukan jika tipe datanya sejenis.

Parameter variadic dapat dikombinasikan dengan parameter biasa, dengan syarat parameter vaiadic-nya harus diposisikan di akhir

*/
