package main

import "fmt"

func main() {
	var poin int
	var hari string

	fmt.Print("Masukkan point Anda: ")
	fmt.Scan(&poin)

	switch poin {
	case 10:
		fmt.Printf("Poin kamu %v, perfect!\n", poin)
	case 9:
		fmt.Printf("Poin kamu %v, awesome!\n", poin)
	case 8:
		fmt.Printf("Poin kamu %v, not bad!\n", poin)
	default:
		fmt.Printf("Poin kamu %v, bad!\n", poin)
	}

	fmt.Print("Masukkan nama hari: ")
	fmt.Scan(&hari)

	switch hari {
	case "senin":
		fmt.Println("monday")
	case "selasa":
		fmt.Println("tuesday")
	case "Rabu":
		fmt.Println("wednesday")
	case "kamis":
		fmt.Println("thursday")
	case "jumat":
		fmt.Println("friday")
	case "sabtu":
		fmt.Println("saturday")
	case "minggu":
		fmt.Println("sunday")
	default:
		fmt.Println("tidak ada nama hari tersebut dalam bahasa inggris")
	}
}

/*
Di Golang, bila sebuah case terpenuhi, tidak akan dilanjutkan ke pengecekan case selanjutnya meskipun tidak ada keyword break.
Konsep tersebu berkebaikan dengan switch pada umumnya
*/
