package main

import (
	"fmt"
	"math"
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

	// fungsi multiple return, ex: menghitung luas dan keliling lingkaran
	var radius float64
	fmt.Printf("Masukkan jari - jari lingkaran: ")
	fmt.Scan(&radius)
	var luas, keliling = luasDanKelilingLingkaran(radius)
	fmt.Printf("Lingkaran dengan jari - jari %v memiliki luas %.2f dan keliling %.2f\n", radius, luas, keliling)

	// fungsi dengan predefined return value
	var sisi float64
	fmt.Printf("Masukkan sisi persegi: ")
	fmt.Scan(&sisi)
	luasPersegi, kelilingPersegi := luasDanKelilingPersegi(sisi)
	fmt.Printf("Persegi dengan sisi %v memiliki luas %v dan keliling %v\n", sisi, luasPersegi, kelilingPersegi)
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

func luasDanKelilingLingkaran(radius float64) (float64, float64) {
	var luas = math.Pi * math.Pow(radius, 2) // hitung luas

	var keliling = 2 * math.Pi * radius // hitung keliling

	return luas, keliling
}

func luasDanKelilingPersegi(sisi float64) (luas float64, keliling float64) {
	luas = math.Pow(sisi, 2)
	keliling = 4 * sisi
	return
}

/*
keyword return sebagai penanda nilai balik dan menghentikan proses dalam blok fungsi di manapun ia dipakai
*/
