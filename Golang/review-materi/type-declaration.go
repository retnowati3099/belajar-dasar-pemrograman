package main

import "fmt"

func main() {
	type noktp string
	var mynoktp noktp = "3403057008990001"

	fmt.Println(mynoktp)
	fmt.Println(noktp("3403051234567990"))
}
