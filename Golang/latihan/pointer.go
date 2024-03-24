package main

import "fmt"

func main() {
	// penerapan pointer
	var num1 int = 5
	var num2 *int = &num1

	fmt.Println(num1)  // 5
	fmt.Println(&num1) // 0xc0000120c0
	fmt.Println(num2)  // 0xc0000120c0
	fmt.Println(*num2) // 5

	// efek perubahan nilai pointer
	num1 = 10
	fmt.Println(num1)
	fmt.Println(&num1)
	fmt.Println(num2)
	fmt.Println(*num2)

	*num2 = 15
	fmt.Println(num1)
	fmt.Println(&num1)
	fmt.Println(num2)
	fmt.Println(*num2)

	num1 = 60
	num2 = &num1
	fmt.Println(num1)
	fmt.Println(&num1)
	fmt.Println(num2)
	fmt.Println(*num2)

	// parameter pointer
	var num3 int = 20
	fmt.Println("Before:", num3)
	change(&num3, 25)
	fmt.Println("After:", num3)
}

func change(original *int, value int) {
	*original = value
}

/*
	- Pointer adalah reference atau alamat memori
	- Variabel pointer berarti variabel yang berisi alamat memori suatu nilai
	- Variabel - variabel yang memiliki reference atau alamat memori yang sama saling berhubungan satu sama lain dan nilainya pasti sama.
	- Ketika ada perubahan nilai, maka akan memberikan efek kepada variabel lain (yang refereninya sama) yaitu nilainya ikut berubah
	- Variabel bertipe pointer ditandai dengan adanya tanda asterik (*) tepat sebelum penulisan tipe data ketika deklarasi
	-Nilai default variabel pointer adalah nil (kosong).
	- Variabel pointer tidak dapat menampung nilai yang bukan pointer, dan sebaliknya variabel biasa tidak bisa menampung nilai pointer
	- Variabel biasa bisa diambil nilai pointernya degan menambahkan tanda ampersand (&) tepat sebelum nama variabel, disebut dengan referencing.
	- Nilai asli variabel pointer juga bisa diambil dengan menambahkan tanda asterik (*) tepat sebelum nama variabel, disebut dengan dereferencing.
	- Ketika salah satu variabel pointer diubah nilainya, sedang ada variabel lain yang memiliki referensi memori yang sama, maka nilai variabel lain tersebut juga akan berubah
*/
