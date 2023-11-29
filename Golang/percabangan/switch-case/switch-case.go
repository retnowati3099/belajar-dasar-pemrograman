package main

import (
	"fmt"
)

func main() {
	var poin int
	var hari string
	var nilai float32
	var num int
	var bulan int

	fmt.Print("Masukkan poin Anda: ")
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
	case "rabu":
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

	// case untuk banyak kondisi
	fmt.Println("Masukkan nilai Anda: ")
	fmt.Scan(&nilai)

	switch nilai {
	case 10:
		fmt.Println("Perfect")
	case 9, 7, 8:
		fmt.Println("Good")
	case 6, 5:
		fmt.Println("Not bad")
	case 4, 3, 2, 1:
		fmt.Println("Bad")
	default:
		fmt.Println("Error")
	}

	// switch gaya if else
	// nilai yang dibandingkan tidak ditulis setelah keyword switch, melainkan akan ditulis langsung dalam bentuk perbandingan dalam keyword case
	fmt.Print("Masukkan sembarang bilangan: ")
	fmt.Scan(&num)

	switch {
	case num > 0:
		fmt.Println(num, "adalah bilangan positif.")
	case (num == 0):
		fmt.Println(num, "adalah nol.")
	default:
		fmt.Println(num, "adalah bilangan negatif")
	}

	// keyword fallthrough
	// keyword fallthrough digunakan untuk memaksa proses pengecekan diteruskan ke satu case selanjutnya dengan tanpa menghiraukan nilai kondisinya, satu case di pengecekan selanjutnya selalu dianggap benar meskipun aslinya adalah salah

	fmt.Print("Inputkan bulan ke- ")
	fmt.Scan(&bulan)
	switch bulan {
	case 1:
		fmt.Println("Januari")
	case 2:
		fmt.Println("Februari")
		fallthrough
	case 3:
		fmt.Println("Maret")
	case 4:
		fmt.Println("April")
	case 5:
		fmt.Println("Mei")
	case 6:
		fmt.Println("Juni")
	case 7:
		fmt.Println("Juli")
	case 8:
		fmt.Println("Agustus")
	case 9:
		fmt.Println("September")
	case 10:
		fmt.Println("Oktober")
	case 11:
		fmt.Println("November")
	case 12:
		fmt.Println("Desember")
	default:
		fmt.Println("Bukan input bulan / salah inputan")
	}
}

/*
Di Golang, bila sebuah case terpenuhi, tidak akan dilanjutkan ke pengecekan case selanjutnya meskipun tidak ada keyword break.
Konsep tersebu berkebaikan dengan switch pada umumnya
Tanda kurung kurawal dapat diterapkan pada keyword case dan defult (opsional).
Bagus jika dipakai pada blok kondisi yang di dalamnya ada banyak statement
*/
