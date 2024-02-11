package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32) // 32768
	fmt.Println(nilai64) // 32768
	fmt.Println(nilai16) // -32768

	var name = "Retno Wati"
	var e = name[0] // type data e: byte atau uint8
	var eString = string(e)

	fmt.Println(name)    // Retno Wati
	fmt.Println(e)       // 82
	fmt.Println(eString) // R

}
