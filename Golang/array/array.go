package main

import "fmt"

func main() {
	// deklarasi array
	var arr_num1 = [3]int{1, 2, 3}
	var arr_num2 = [...]int{1, 4, 9}
	arr_num3 := [3]int{2, 4, 6}
	arr_num4 := [...]int{1, 3, 5}
	var name = [3]string{"wildan", "pandu", "aska"}

	// cetak elemen array
	fmt.Println(arr_num1)
	fmt.Println(arr_num2)
	fmt.Println(arr_num3)
	fmt.Println(arr_num4)
	fmt.Println(name)

	// akses dan cetak elemen tetentu dari array
	fmt.Println(name[0])
	fmt.Println(name[2])

	// ubah elemen tertentu di dalam array
	name[1] = "khresna"
	fmt.Println(name) // cetak array name setelah terjadi perubahan

	// inisialisasi array <--jika array tidak diinisialisasi, maka array tersebut akan diberi nilai default, 0 untuk int dan "" untuk string
	empty := [3]int{}       // tidak diinisialisasi
	partial := [3]int{1, 2} // diinisialisasi sebagian atau parsial
	full := [3]int{1, 2, 3} // diinisialisasi lengkap
	fmt.Println(empty)
	fmt.Println(partial)
	fmt.Println(full)

	// inisialisasai spesifik elemen
	arr_num5 := [3]int{1: 10, 2: 5}
	fmt.Println(arr_num5)

	//panjang array menggunkaan fungsi len()
	fmt.Println(len(name))
	fmt.Println(len(arr_num3))

	var fruits [3]string // deklarasi array
	fruits[0] = "pineapple"
	fruits[1] = "apple"
	fruits[2] = "orange"
	fmt.Println(fruits[0], fruits[1], fruits[2])
	fmt.Println("Jumlah elemen di dalam array fruits adalah", len(fruits))

	var flowers [3]string

	// inisialisasi nilai array dengan gaya vertikal
	flowers = [3]string{"sunflower", "rose", "jasmine"}
	fmt.Println(flowers)

	// inisialisasi nilai array dengan gaya horisontal
	flowers = [3]string{
		"dandelion",
		"tulip",
		"hibiscus",
	}
	fmt.Println(flowers)

	var matriks1 = [2][2]int{{1, 2}, {5, 6}}
	var matriks2 = [2][2]int{[2]int{1, 2}, [2]int{5, 6}}
	fmt.Println(matriks1)
	fmt.Println(matriks2)

	// perulangan elemen array menggunakan keyword for
	var animals = [3]string{"cow", "goat", "sheep"}
	for i := 0; i < len(animals); i++ {
		fmt.Printf("Elemen %d: %s\n", i, animals[i])
	}

	// perulangan elmen array menggunakan keyword for - range
	for i, animal := range animals {
		fmt.Printf("Elemen %d: %s\n", i, animal)
	}

	// penggunaan variabel underscore dalam for - range
	for _, animal := range animals {
		fmt.Printf("Nama binatang: %s\n", animal)
	}
}

// Deklarasi Array
// 1. Menggunakan keyword var
// var array_name = [length]datatype{values} <-- panjang didefinisikan
// atau
// var array_name = [...]datatype{values} <-- panjang tidak didefinisikan, compiler memutuskan panjang array berdasarkan banyak nilainya
// atau
// var array_name [length]datatype
// 2. Menggunakan tanda :=
// array_name := [length]datatype{values} atau
// array_name := [...]datatype{values}

/*
Inisialisasi
Inisialisasi array dengan gaya vertikal, tanda koma wajib dituliskan setelah elemen, termasuk elemen terakhir. Jika tidak, maka akan muncul error
*/
