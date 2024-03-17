package main

import "fmt"

// deklarasi konstanta di luar fungsi
const PI = 3.14
const (
	yesterday string = "Minggu"
	kemarin          // manifest typing
	isToday3  = true
)

func main() {
	fmt.Printf("Phi: %.2f\n", PI)

	// penggunaan konstanta
	const firstName string = "Retno"
	fmt.Print("Hallo ", firstName, "!\n")

	const lastName = "Wati"
	fmt.Print("Nice to meet you ", lastName, "!\n")

	// penulisan beda, output sama
	fmt.Println("Retno Wati")
	fmt.Println(firstName, lastName)
	fmt.Print("Retno Wati\n")
	fmt.Print("Retno ", "Wati\n")
	fmt.Print("Retno", " Wati\n")
	fmt.Print(firstName, " ", lastName, "\n")
	fmt.Printf("%s %s\n", firstName, lastName)

	// deklarasi multi konstanta di dalam fungsi
	// deklarasi konstanta dengan tipe data dan nilai yang berbeda
	const (
		square         = "kotak" // type inference
		isToday  bool  = true    // manifest typing
		numeric  uint8 = 1       // manifest typing
		floatNum       = 2.2     // type inference
	)

	// deklarasi konstanta dengan tipe dan nilai yang sama
	const (
		a = "konstanta"
		b // type inference
	)

	// gabungan
	const (
		today    string = "Minggu"
		sekarang        // manifest typing
		isToday2 = true
	)

	// deklarasi multiple konstanta dalam satu baris
	const satu, dua = 1, 2
	const three, four string = "tiga", "empat"
}

/*
	- Konstanta adalah jenis variabel yang nilainya tidak bisa diubah.
	- Inisialisasi nilai konstanta hanya dilakukan sekali di awal, setelah itu variabel tidak bisa diubah nilainya
	- Cara penerapan konstanta sama seperti deklarasi variabel biasa, selebihnya tinggal ganti keyword var dengan const
	- Teknik type inference bisa diterapkan pada konstanta, caranya dengan menghilangkan tipe data pada saat deklarasi

	- Fungsi fmt.Print() memiliki peran yang sama seperti fungsi fmt.Println(), bedanya fmt.Print() tidak menghasilkan baris baru di akhir outputnya dan nilai parameter - parameter yang dimasukkan ke dalam fungsi fmt.Print() digabungkan tanpa pemisah.
	- Fungsi fmt.Println() nilai parameternya digabung menggunakan penghubung spasi.

	- Ketika tipe data dan nilai tidak dituliskan dalam deklarasi konstanta, maka tipe data dan nilai yang dipergunakan sama seperti konstanta yang dideklarasikan di atasnya

	- Nama konstanta biasanya ditulis dalam huruf besar (UPPERCASE)
	- Konstanta dapat dideklarasi baik di dalam maupun di luar fungsi
*/
