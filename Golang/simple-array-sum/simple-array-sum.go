package main

import "fmt"

func main() {
	var arrLen int
	fmt.Printf("Panjang array = ")
	fmt.Scan(&arrLen)

	var total = arrSum(arrLen)
	fmt.Printf("Jumlah elemen array adalah %v\n", total)
}

func arrSum(arrLen int) int {
	var arr = make([]int, arrLen)
	total := 0

	for i := 0; i < arrLen; i++ {
		fmt.Printf("Elemen ke-%v = ", i+1)
		fmt.Scan(&arr[i])
		total += arr[i]
	}
	return total
}
