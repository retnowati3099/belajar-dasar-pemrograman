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
*/
