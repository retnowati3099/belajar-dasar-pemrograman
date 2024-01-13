package main

import "fmt"

// deklarasi struct
type student struct {
	name  string
	grade float64
}

func main() {
	var s1 student
	s1.name ="Retno Wati"
	s1.grade = 3.53

	fmt.Println("Name: ", s1.name)
	fmt.Println("Grade: ", s1.grade)
}

/*
- Go tidak memiliki class yang ada di bahasa - bahasa strict OOP lain, tapi Go memiliki tipe data struktur yang disebut struct
- Struct adalah kumpulan definisi variabel (atau properti) dan atau fungsi (atau method), yang dibungkus sebagai tipe data baru dengan nama tertentu.
- Properti dalam struct, tipe datanya bisa bervariasi. Mirip dengan map, hanya saja key-nya sudah didefinisikan di awal, dan tipe data tiap itemnya bisa berbeda
- Dari sebuah struct, kita bisa buat variabel baru, yang memiliki atribut sesuai skema struct tersebut
- Keyword type digunakan untuk deklarasi struct
*/
