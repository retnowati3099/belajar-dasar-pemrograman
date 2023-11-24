package main

import "fmt"

func main() {
	var nem int

	fmt.Print("Masukkan nem kamu: ")
	fmt.Scan(&nem)

	if mean := nem / 4; mean == 100 {
		fmt.Printf("Rata - rata ujian kamu %v, perfect !\n", mean)
	} else if mean >= 90 {
		fmt.Printf("Rata - rata ujian kamu %v, good !\n", mean)
	} else if mean >= 70 {
		fmt.Printf("Rata - rata ujian kamu %v, not bad !\n", mean)
	} else {
		fmt.Printf("Rata - rata ujian kamu %v, bad !\n", mean)
	}
}

/*
variabel temporary adalah variabel yang hanya bisa digunakan pada blok seleksi kondisi dimana ia ditempatkan saja.
Berikut adalah beberapa manfaat penggunaan variabel temporary.
1. Scope atau cakupan variabelnya jelas
2. Kode menjadi lebih rapi
3. Ketika nilai variabel didapat dari sebuah komputasi, perhitungan tidak perlu dilakukan di dalam blok masing - masing kode

Deklarasi variabel temporary hanya bisa dilakukan lewat metode type inference, yang menggunakan tanda :=. Peggunaan keyword var tidak diperbolehkan karena akan menyebabkan error
*/
