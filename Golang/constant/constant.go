package main

import (
	"fmt"
)

const PI float32 = 3.14         // typed constants
const PERCEPATAN_GRAVITASI = 10 // untyped constants

// deklarasi lebih dari satu konstanta
const (
	A int = 1
	B     = 3.14
	C     = "Hello"
)

func main() {
	fmt.Println(PI)
	fmt.Println(PERCEPATAN_GRAVITASI)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}

// Penulisan sintaks konstanta
// const CONATNAME type = value

// Aturan umum untuk konstata
// 1. Nama konstanta mengikuti aturan penamaan pada variabel
// 2. Nama konstanta biasanya ditulis dalam huruf besar (untuk memudahkan identifikasi dan membedakannya dengan variabel)
// 3. Konstanta dapat dideklarasi baik di dalam ataupun di luar fungsi
// 4. Konstanta tidak dapat diubah dan bersifat read only

// Tipe konstanta:
// 1. Typed constants -> dideklarasi dengan tipe yang terdefinisi
// 2. Untyped constants -> dideklarasi tanpa menyertakan tipenya. berarti tipenya mengacu dari nilai (compilermemutuskan tipe konstanta berdasarkan pada nilai)
