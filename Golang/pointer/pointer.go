package main

import "fmt"

func main() {
	var age *int
	var name *string
	fmt.Println(age)  // <nil>
	fmt.Println(name) // <nil>

	var numA int = 5
	var numB *int = &numA
	fmt.Println("Value of numA: ", numA)
	fmt.Println("Address of numA: ", &numA)
	fmt.Println("Value of numB: ", *numB)
	fmt.Println("Address of numB: ", numB)

	numA = 7
	fmt.Println("Value of numA: ", numA)
	fmt.Println("Address of numA: ", &numA)
	fmt.Println("Value of numB: ", *numB)
	fmt.Println("Address of numB: ", numB)

	// parameter pointer
	var numC = 5
	fmt.Println("Before: ", numC)

	change(&numC, 10)
	fmt.Println("After: ", numC)

	// menampilkan tulisan "Praktikum Pemrograman" menggunakan variabel pointer
	var sentence1 string = "Praktikum Pemrograman"
	var sentence2 *string = &sentence1
	fmt.Println("Sentence2 : ", *sentence2)

	// program dengan fungsi dan variabel pointer untuk menukar dua isi variabel biasa
	var a, b int

	fmt.Printf("Masukkan nilai sembarang bilangan ke-1 : ")
	fmt.Scan(&a)
	fmt.Printf("Masukkan nilai sembarang bilangan ke-2 : ")
	fmt.Scan(&b)

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

	tukar(&a, &b)

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
}

func change(num *int, value int) {
	*num = value
}

func tukar(num1, num2 *int) {
	var temp int
	temp = *num1
	*num1 = *num2
	*num2 = temp
}

/*
- Pointer adalah reference atau alamat memori
- Variabel pointer berarti variabel yang berisi alamat memori suatu nilai
- Misal sebuah variabel bertipe integer memiliki nilai 4, maka yang dimaksud pointer adalah alamat memori di mana nilai 4 disimpan, bukan nilai 4 itu sendiri
- Variabel - variabel yang memiliki reference atau alamat pointer yang sama, saling berhubungan satu sama lain dan nilainya pasti sama. Ketika ada perubahan nilai, maka akan memberikan efek kepada variabel lain (yang referensi-nya sama) yaitu nilainya ikut berubah.
- Variabel bertipe pointer ditandai dengan adanya tanda asterik (*) tepat sebelum penulisan tipe data ketika deklarasi.
- Nilai default variabel pointer adalah nil (kosong).
- Variabel pointer tidak bisa menampung nilai yang bukan pointer, dan sebaliknya variabel biasa tidak bisa menampung nilai pointer.
- Variabel biasa bisa diambil nilai pointernya dengan cara menambahkan tanda ampersand (&) tepat sebelum nama variabel. Metode ini disebut dengan referencing.
- Nilai asli variabel pointer juga bisa diambil dengan cara menambahkan tanda asterik (*) tepat sebelum nama variabel. Metode ini disebut deferencing.
- Penerapan pointer sebagai parameter kurang lebih sama dengan cara mendeklarasikan parameter sebagai pointer
*/
