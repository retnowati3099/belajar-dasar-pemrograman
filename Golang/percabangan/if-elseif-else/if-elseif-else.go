package main

import "fmt"

func main() {
	var num int

	fmt.Print("Masukkan sembarang bilangan: ")
	fmt.Scan(&num)

	if num > 0 {
		fmt.Println(num, "adalah bilangan positif.")
	} else if num == 0 {
		fmt.Println(num, "adalah nol.")
	} else {
		fmt.Println(num, "adalah bilangan negatif")
	}
}

// penggunaan if-else di golang sama seperti bahasa pemrograman lain, yang membedakan adalah tanda kurungnya, di golang tidak perlu ditulis
// Kurung kurawal harus dituliskan meski isinya hanya satu blog statement
