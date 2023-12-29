package main

import "fmt"

func main(){
	var a, b int
	fmt.Printf("a = ")
	fmt.Scan(&a)
	fmt.Printf("b = ")
	fmt.Scan(&b)

	fmt.Println(solveMeFirst(a, b))
}

func solveMeFirst(a, b int) int {
	return a + b
}