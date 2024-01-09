package main

import "fmt"

func main() {
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				min, max = e, e
			case e < min:
				min = e
			case e > max:
				max = e
			}
		}
		return min, max
	}

	var lenNums int
	fmt.Printf("Masukkan panjang array: ")
	fmt.Scan(&lenNums)
	var nums = make([]int, lenNums)
	for i := 0; i < lenNums; i++ {
		fmt.Printf("Elemen ke %v: ", i+1)
		fmt.Scan(&nums[i])
	}
	var min, max = getMinMax(nums)
	fmt.Printf("Data %v\nMin: %v\nMax: %v\n", nums, min, max)
}

/*
- Closure adalah fungsi yang bisa disimpan dalam variabel.
- Variabel yang menyimpan closure memiliki sifat seperti fungsi yang disimpannya
- Konsep closure memungkinkan untuk bisa membuat fungsi di dalam fungsi, atau bahkan membuat fungsi yang mengembalikan fungsi
- Closure merupakan anonymous function atau fungsi tanpa nama.
- Closure biasa dimanfaatkan untuk membungkus suatu proses yang hanya dipakai sekali atau dipakai pada blok tertentu saja
*/
