package main

import "fmt"

func main() {
	var num1, num2, jumlah int

	fmt.Print("Masukkan bilangan pertama: ")
	fmt.Scan(&num1)
	fmt.Print("Masukkan bilangan kedua: ")
	fmt.Scan(&num2)

	jumlah = num1 + num2

	println(num1, "+", num2, "=", jumlah)

}
