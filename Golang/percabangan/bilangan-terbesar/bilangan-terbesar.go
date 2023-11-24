package main

import "fmt"

func main() {
	var bil1, bil2, bil3 int

	fmt.Print("Inputkan bilangan 1: ")
	fmt.Scan(&bil1)
	fmt.Print("Inputkan bilangan 2: ")
	fmt.Scan(&bil2)
	fmt.Print("Inputkan bilangan 3: ")
	fmt.Scan(&bil3)

	if bil1 >= bil2 && bil1 >= bil3 {
		fmt.Printf("Bilangan terbesar adalah %d.\n", bil1)
	} else if bil2 >= bil1 && bil2 >= bil3 {
		fmt.Printf("Bilangan terbesar adalah %d.\n", bil2)
	} else {
		fmt.Printf("Bilangan terbesar adalah %d.\n", bil3)
	}
}
