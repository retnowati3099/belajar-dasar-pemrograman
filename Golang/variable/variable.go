package main

import (
	"fmt"
)

var width int = 36
var height int

func main() {
	var firstName string = "Retno"
	var lastName = "Wati"
	fullName := "Retno Wati"

	var myName string
	var age int
	var weight float32
	var isFun bool

	var student1 string
	student1 = "Aska"

	height = 142

	var x, y, z int = 1, 2, 3
	var a, b = "Afriska", 150
	c, d := "Retno", 142

	var (
		e int
		f int    = 1
		g string = "hello"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(fullName)
	fmt.Print(firstName + " " + lastName + "\n")
	fmt.Println(firstName + " " + lastName)
	fmt.Printf(firstName + " " + lastName + "\n")
	fmt.Printf("%s %s \n", firstName, lastName)

	fmt.Println(myName)
	fmt.Println(weight)
	fmt.Println(age)
	fmt.Println(isFun)

	fmt.Println(student1)

	fmt.Println(width)
	fmt.Println(height)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}

/*
Aturan penamaan variabel:
1. Nama variabel harus dimulai dengan huruf atau underscore(_)
2. Nama variabel tidak dapat dimulai dengan digit
3. Nama variabel hanya dapat berisi karakter apha-numeric dan underscore (a-z, A-Z, 0-9, dan _)
4. Nama variabel case-sensitive (age, Age, dan AGE adalah tiga variabel yang berbeda)
5. Tidak ada batasan panjang nama variabel
6. Nama variabel tidak dapat berisi spasi
7. Yang termasuk keyword dalam bahasa Go tidak boleh dijadikan sebagai nama variabel
*/
