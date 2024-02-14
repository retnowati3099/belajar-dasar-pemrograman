package main

import "fmt"

func main() {
	var a = 20
	var b = 10

	// operator aritmatika
	var c = a + b // penjumlahan
	var d = a - b // pengurangan
	var e = a * b // perkalian
	var f = a / b // pembagian
	var g = a % b // modulo atau sisa pembagian

	fmt.Println("a = ", a) // 20
	fmt.Println("b = ", b) // 10
	fmt.Println("c = ", c) // 30
	fmt.Println("d = ", d) // 10
	fmt.Println("e = ", e) // 200
	fmt.Println("f = ", f) // 2
	fmt.Println("g = ", g) // 0

	// augmented assignments
	a += 10 // a = a + 10
	fmt.Println("a = ", a)
	a -= 10 // a = a - 10
	fmt.Println("a = ", a)
	a *= 10 // a = a * 10
	fmt.Println("a = ", a)
	a /= 10 // a = a / 10
	fmt.Println("a = ", a)
	a %= 10 // a = a % 10
	fmt.Println("a = ", a)

	// operator unary
	var h = 1
	h++ // increment -> h = h + 1
	fmt.Println(h)
	h-- // decrement -> h = h - 1
	fmt.Println(h)

	// operator perbandingan atau relasi
	var name1 = "Retno"
	var name2 = "Retno"
	var result bool = name1 == name2
	fmt.Println(result)
	fmt.Println(name1 != name2)
	fmt.Println(3 >= 4)
	fmt.Println(3 > 4)
	fmt.Println(3 <= 4)
	fmt.Println(3 < 4)

	// operator logika atau boolean
	var nilaiAkhir = 90
	var absensi = 80
	var lulusNilaiAKhir bool = nilaiAkhir > 80    // true
	var lulusPresensi bool = absensi > 80         // false
	var lulus = lulusNilaiAKhir && lulusPresensi  // true and false = false
	var lulus2 = lulusNilaiAKhir || lulusPresensi // true or false = true
	fmt.Println(lulus)                            // false
	fmt.Println(lulus2)

}
