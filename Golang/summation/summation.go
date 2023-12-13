package main

import "fmt"

func main() {
	var panjang int
	var sum int = 0

	fmt.Print("Masukkan jumlah angka yang ingin dijumlahkan: ")
	fmt.Scan(&panjang)
	
	temp := make([]int, panjang)

	for i := 0; i < panjang; i++ {
		fmt.Printf("Angka ke-%d: ", (i + 1))
		fmt.Scan(&temp[i])
		sum += temp[i]
	}
	fmt.Println(sum)
}
