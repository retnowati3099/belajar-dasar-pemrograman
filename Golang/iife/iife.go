package main

import "fmt"

func main() {
	var arrLen int
	fmt.Printf("Masukkan panjang array: ")
	fmt.Scan(&arrLen)

	var arr = make([]int, arrLen)
	for i := 0; i < arrLen; i++ {
		fmt.Printf("Elemen ke %v: ", i+1)
		fmt.Scan(&arr[i])
	}
	// fmt.Println(arr)

	var newArr = func(min int) []int {
		var result []int
		for _, e := range arr {
			if e < min {
				continue
			}
			result = append(result, e)
		}
		return result
	}(2)
	fmt.Println("Original array: ", arr)
	fmt.Println("Filtered array: ", newArr)
}

/*
- IIFE (Immediatly-Invoked Funtion Expression) meupakan jenis closure yang dieksekusi langsung pada saat deklarasinya. Biasa digunakan untuk membungkus proses yang hanya dilakukan sekali, bisa mengembalikan nilai, bisa juga tidak.
- Ciri khas IIFE adalah adanya kurung parameter tepat setelah deklarasi closure berakhir. Jika ada parameter, bisa juga dituliskan dalam kurung parameternya.
- Pada contoh di atas, IIFE menhasilkan nilai balik yang kemudian ditampung newArr. Yang ditampung adalah  nilai kembaliannya bukan body fungsi atau closure.
- Closure bisa juga ditulis dengna gaya manifest typing, caranya dengan menuliskan skema closure-nya sebagai tipe data. Contohnya sebagai berikut:
var closure (func (string, int, []string) int)
closure = func (a string, b int, c []string) int {
	// ...
}
*/
