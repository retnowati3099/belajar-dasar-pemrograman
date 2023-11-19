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
}

// Deklarasi Array
// 1. Menggunakan keyword var
// var array_name = [length]datatype{values} <-- panjang didefinisikan
// atau
// var array_name = [...]datatype{values} <-- panjang tidak didefinisikan, compiler memutuskan panjang array berdasarkan banyak nilainya
// 2. Menggunakan tanda :=
// array_name := [length]datatype{values} atau
// array_name := [...]datatype{values}
