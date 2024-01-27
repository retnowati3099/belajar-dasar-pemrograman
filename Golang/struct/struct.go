package main

import "fmt"

// deklarasi struct -> struct student dideklarasikan memiliki dua properti yaitu name dan grade
type student struct {
	name  string
	grade float64
}

func main() {
	var s1 student
	s1.name = "Retno Wati"
	s1.grade = 3.53

	fmt.Println("Name: ", s1.name)
	fmt.Println("Grade: ", s1.grade)

	var s2 = student{}
	s2.name = "Afifah Nur Oktavia"
	s2.grade = 4.00

	fmt.Println("Name: ", s2.name)
	fmt.Println("Grade: ", s2.grade)

	var s3 = student{"Latri Nur Fadhilah", 3.99}

	fmt.Println("Name: ", s3.name)
	fmt.Println("Grade: ", s3.grade)

	var s4 = student{name: "Tomi"}

	fmt.Println("Name: ", s4.name)
	fmt.Println("Grade: ", s4.grade)

	var s5 = student{name:"Afriska Yusuf Widyamto", grade: 3.00}
	var s6 *student = &s5
	fmt.Println("Name: ", s5.name)
	fmt.Println("Name: ", s6.name)

	s6.name = "Afriska Y. W."
	fmt.Println("Name: ", s5.name)
	fmt.Println("Name: ", s6.name)
}

/*
- Go tidak memiliki class yang ada di bahasa - bahasa strict OOP lain, tapi Go memiliki tipe data struktur yang disebut struct
- Struct adalah kumpulan definisi variabel (atau properti) dan atau fungsi (atau method), yang dibungkus sebagai tipe data baru dengan nama tertentu.
- Properti dalam struct, tipe datanya bisa bervariasi. Mirip dengan map, hanya saja key-nya sudah didefinisikan di awal, dan tipe data tiap itemnya bisa berbeda
- Dari sebuah struct, kita bisa buat variabel baru, yang memiliki atribut sesuai skema struct tersebut
- Keyword type digunakan untuk deklarasi struct.
- Cara membuat variabel objek sama seperti pebuatan variabel biasa, yaitu dengan menulis nama variabel diikuti nama struct.
- Semua property variabel objek pada awalnya memiliki zero value sesuai tipe datanya.
- Property variabel objek bisa diakses nilainya menggunakan notasi titik.
- Cara inisialisasivariabel objek adalah dengan menambahkan kurung kurawal setelah nama struct. Nilai masing - masing property bisa diisi pada saat inisialisasi.
- Objek yang dibuat dari tipe struct bisa diambil nilai pointernya dan bisa disimpan pada variabel objek yang bertipe struct pointer.
*/
