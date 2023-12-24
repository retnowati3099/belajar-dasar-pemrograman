package main

import (
	"fmt"
	"strings"
)

func main() {
	var names = []string{"Retno", "Wati"}
	printMessage("Hallo", names)

	var word string
	fmt.Printf("Masukkan sembarang kata: ")
	fmt.Scan(&word)
	var pw = printWord(word)
	fmt.Println("Hello", pw)

	// konversi desimal ke biner
	var dec int
	fmt.Printf("Masukkan sembarang bilangan desimal: ")
	fmt.Scan(&dec)
	var hasil = decimalToBiner(dec)
	fmt.Println(hasil)
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

func printWord(word string) string {
	return word
}

func decimalToBiner(dec int) []int {
	var arrBiner []int
	var sisa int
	for dec > 0 {
		sisa = dec % 2
		arrBiner = append(arrBiner, sisa)
		dec /= 2
	}
	n := len(arrBiner)
	for i := 0; i < n/2; i++ {
		arrBiner[i], arrBiner[n-i-1] = arrBiner[n-i-1], arrBiner[i]
	}
	return arrBiner
}

/*
keyword return sebagai penanda nilai balik dan menghentikan proses dalam blok fungsi di manapun ia dipakai
*/
