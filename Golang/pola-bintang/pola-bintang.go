package main

import "fmt"

func main() {
	var baris int

	fmt.Print("Masukkan banyak baris: ")
	fmt.Scan(&baris)

	for i := 1; i <= baris; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	fmt.Println("===========")

	var bintang string = ""
	for i := 1; i <= baris; i++ {
		bintang += "* "
		fmt.Println(bintang)
	}

	fmt.Println("===========")

	for i := 1; i <= baris; i++ {
		// spasi
		for j := (baris - 1); j >= i; j-- {
			fmt.Print(" ")
		}

		// bintang
		for k := 1; k <= i; k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
