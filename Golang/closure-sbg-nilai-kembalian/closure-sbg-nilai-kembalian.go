package main

import "fmt"

func main() {
	var max int = 10
	var arrLen int
	fmt.Printf("Masukkan panjang array: ")
	fmt.Scan(&arrLen)

	var arr = make([]int, arrLen)
	for i := 0; i < arrLen; i++ {
		fmt.Printf("Elemen-%v: ", i+1)
		fmt.Scan(&arr[i])
	}

	var howMany, getNumbers = findMax(arr, max)
	var theNumbers = getNumbers()

	fmt.Println("Numbers\t: ", arr)
	fmt.Println("Find\t:", max)
	fmt.Println("Found\t:", howMany)
	fmt.Println("Value\t:", theNumbers)
}

func findMax(arr []int, max int) (int, func() []int) {
	var res []int
	for _, e := range arr {
		if e <= max {
			res = append(res, e)
		}
	}

	var getNumbers = func() []int {
		return res
	}

	return len(res), getNumbers
}
