package main

import "fmt"

func main() {
	// deklarasi variabel beserta tipe datanya (manifest typing)
	var firstName string = "Retno"
	var lastName string
	lastName = "Wati"

	// output sama, penulisan beda
	fmt.Printf("Hallo %s %s!\n", firstName, lastName)
	fmt.Printf("Hallo Retno Wati!\n")
	fmt.Println("Hallo", firstName, lastName+"!")

	// deklarasi variabel tanpa tipe data (type inference)
	fullName := "Retno Wati"
	fmt.Printf("Hello %s!\n", fullName)

	var student1 = "Afriska" // menggunakan var, tanpa tipe data, menggunakan perantara "="
	student2 := "Retno"      // tanpa var, tanpa tipe data, menggunakan perantara ":="
	fmt.Printf("Student 1: %s\nStudent 2: %s\n", student1, student2)

	student2 = "Retno Wati"
	fmt.Printf("Student 2: %s\n", student2)

	// deklarasi multivariabel
	var first, second, third string = "satu", "dua", "tiga" // cara 1
	var fourth, fifth, sixth string                         // cara 2
	fourth, fifth, sixth = "empat", "lima", "enam"
	seventh, eight, ninth := "tujuh", "delapan", "sembilan" // cara 3
	println(first, second, third, fourth, fifth, sixth, seventh, eight, ninth)

	// deklarasi multi variabel, tipe data berbeda, metode type inference
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"
	fmt.Println(one, isFriday, twoPointTwo, say)

	// variabel underscore (_)
	_ = "Belajar Golang"
	_ = "Golang itu mudah"
	name, _ := "Retno", "Wati"
	fmt.Println(name)

	// deklarasi variabel menggunakan keyword new
	employee := new(string) // variabel employee menampung data bertipe pointer string
	fmt.Println(employee)   // output berupa alamat memori (dalam bentuk notasi heksadesimal)
	fmt.Println(*employee)  // tanda asterik (*) digunakan untuk men-dereference variabel agar nilai aslinya bisa ditampilkan

	// deklarasi variabel dengan keyword make
	// tidak akan dicontohkan di sini
}

/*
	- Go mengadopsi dua jenis penulisan variabel, yaitu dituliskan tipe datanya (manifest typing) dan juga tidak(type inference).
	- Keyword var digunakan untuk deklarasi variabel (membuat variabel baru).
	- Nilai variabel bisa diisi langsung pada saat deklarasi variabel.

	- Fungsi fmt.Printf() digunakan untuk menampilkan output dalam bentuk tertentu.
	- Penggunaannya sama dengan fmt.Println(), hanya saja struktur outputnya didefinisikan di awal
	- Karakter %s akan diganti dengan data string yang berada di parameter.
	- Fungsi fmt.Printf() tidak menghasilkan baris baru di akhir teks, sehingga digunakan literal newline \n untuk memunculkan baris baru di akhir.
	- Tanda + bila ditempatkan di antara string, fungsinnya adalah untuk penggabungan string (string concatenation)
	- Fungsi fmt.Println() akan secara otomatis menghasilkan baris baru (new line) di akhir

	- Go mengadopsi konsep type inference yaitu metode deklarasi variabel yang tipe datanya ditentukan oleh tipe data nilainya (keyword var dan tipe data tidak perlu ditulis).
	- Pada metode type inference, operand = harus diganti dengan := dan keyword var dihilangkan.
	- Diperbolehkan untuk tetap menggunakan keyword var pada saat deklarasi meskipun tanpa menuliskan tipe data dengan ketentuan tidak menggunakan tanda :=, melainkan tetap menggunakan =.
	- Tanda := hanya digunakan sekali di awal pada saat deklarasi, untuk assignment nilai selanjutnya harus menggunakan tanda =.
	- deklarasi menggunakan := hanya bisa dilakukan di dalam blok fungsi

	- Go mendukung metode deklaasi banyak variabel secara bersamaan, caranya dengan menuliskan variabel variabelnya dengan pembatas tanda koma.
	- Pengisian nilai juga bisa dilakukan bersamaan pada saat deklarasi, caranya dengan menuliskan nilai masing - masing variabel berurutan sesuai variabelnya dengan pembatas koma
	- Teknik type inference, deklarasi multi variabel bisa dilakukan untuk variabel - variabel yang tipe data satu sama lainnya berbeda.

	- Di Go tidak boleh ada satu pun variabel yang menganggur, semua variabel yang dideklarasikan harus digunakan.
	- Jika terdapat variabel yang dideklarasikan tetapi tidak digunakan, error akan muncul pada saat kompilasi dan program tidak akan bisa dirun
	- Underscore (_) adalah reserved variabel yang bisa dimanfaatkan untuk menamung nilai yang tidak dipakai (keranjang sampah)
	- variabel underscore adalah predefined, tidak perlu menggunakan := untuk pengisian nilai, cukup dengan = saja.
	- Khusus untuk pengisian nilai multi variabel yang digunakan menggunakan metode type inference, boleh di dalamnya terdapat variabel underscore.
	- Biasanya variabel underscore sering digunakan untuk menampung nilai balik fungsi yang tidak digunakan
	- Isi variabel underscore tidak dapat ditampilkan, data yang sudah masuk variabel tersebut akan hilang.

	- Keyword new digunakan untuk membuat variabel pointer dengan tipe data tertentu.
	- Nilai data defaultnya akan menyesuaikan tipe datanya.

	- Keyword make hanya dapat digunakan untuk pembuatan beberapa jenis variabel saja, yaitu slice, map, dan channel
*/
