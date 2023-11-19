package main

import (
	"fmt"
)

func main() {
	var word1, word2 string = "Hello", "World"
	fmt.Print(word1)
	fmt.Print(word2)

	// cetak argumen dalam baris baru, butuh \n
	fmt.Print(word1, "\n")
	fmt.Print(word2, "\n")

	// menggunakan satu fungsi Print() untuk pencetakan banyak variabel
	fmt.Print(word1, "\n", word2)

	// menambahkan spasi di antara argumen string, butuh ""
	fmt.Print(word1, " ", word2)

	// Print() menyertakan spasi diantara argumen jika keduanya bukan string
	var num1, num2 = 24, 25
	fmt.Print(num1, num2)

	fmt.Println(word1, word2)

	var num int = 24
	var name string = "Retno"
	fmt.Printf("name memiliki nilai %v dan tipe %T\n", name, name)
	fmt.Printf("num memiliki nilai %v dan tipe %T\n", num, num)
}

// Terdapat 3 fungsi untuk menampilkan teks, yaitu Print(), Println, dan Printf()
// Fungsi Print() menampilkan argumen dengan format default
// Fungsi Println() sama dengan Print() tetapi spasi ditambahkan di antara argumen, dan baris baru ditambahkan di akhir
// Fungsi Printf() memformat argumen berdasarkan pada formmating verb yang diberikan kemudian mencetak mereka
// %v digunakan untuk mencetak nilai dari argumen
// %T diguanakn untuk mencetak tipe dari argumen
