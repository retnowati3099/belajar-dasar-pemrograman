package main

import "fmt"

func main() {
	var arr_num1 = [3]int{1, 2, 3}
	var arr_num2 = [...]int{1, 4, 9}

	fmt.Println(arr_num1)
	fmt.Println(arr_num2)
}

// Deklarasi Array
// 1. Menggunakan keyword var
// var array_name = [length]datatype{values} <-- panjang didefinisikan
// atau
// var array_name = [...]datatype{values} <-- panjang tidak didefinisikan
