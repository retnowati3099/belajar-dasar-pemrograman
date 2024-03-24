package main

import (
	"fmt"
	"strings"
)

type student struct {
	name  string
	grade float32
}

// fungsi sayHello dideklarasikan sebagai method milik struct student
func (s student) sayHello() {
	fmt.Println("Hallo", s.name)
}

func (s student) getName(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

func main() {
	var s1 = student{"Retno Wati", 24}
	s1.sayHello()

	fmt.Println("nama panggilan :", s1.getName(1))
	var name = s1.getName(2)
	fmt.Println("nama panggilan :", name)
}

/*
	- Method adalah fungsi yang menempel pada type (bisa struct atau tipe data lainnya).
	- Method bisa diakses mellui variabel objek.
	- keunggulan method dibbanding fungsi biasa adalaj memiliki akses ke properti struct hingga ke level private. Selain itu, dengan menggunakan method sebuah proses bisa di-encapsulasi dengan baik.
	- Ketika deklarasi, ditentukan juga siapa pemilik method tersebut. Jadi, cara deklarasi method sama seperti fungsi, hanya saja perlu ditambahkan deklarasi variabel objek di sela - sela keyword func dan nama fungsi. Struct yang digunakan akan menjadi pemilik method.
	- Cara pengaksesan method sama seperti pengaksesam properti berupa variabel.
	- Method memiliki  sifat yang sama persis dengan fungsi biasa, seperti bisa berparameter, memiliki nilai  balik, dll.
	- Dari segi sintaks pembedanya hanya ketika pengaksesan dan deklarasi

*/
