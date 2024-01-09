package main

import "fmt"

func main() {
	var bil int
	fmt.Printf("Masukkan sembarang bilangan = ")
	fmt.Scan(&bil)

	var hasil = hitungFaktorial(bil)
	fmt.Printf("%v! adalah %v\n", bil, hasil)
}

func hitungFaktorial(bil int) int {
	if bil == 0 {
		return 1
	} else {
		return bil * hitungFaktorial(bil-1)
	}
}
