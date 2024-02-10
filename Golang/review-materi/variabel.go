package main

import "fmt"

// deklarasi banyak variabel di luar fungsi main
var (
	firstName = "Retno"
	lastName  = "Wati"
)

func main() {
	var name string
	name = "Retno Wati"
	fmt.Println(name)

	name = "Afriska Yusuf Widyamto"
	fmt.Println(name)

	fullName := "Queena Putri Askawati"
	fmt.Println(fullName)

	fullName = "Aska Putra Pratama"
	fmt.Println(fullName)

	fmt.Println(firstName)
	fmt.Println(lastName)

	// deklarasi banyak variabel di dalam fungsi main
	var (
		fruit1 = "pineapple"
		fruit2 = "banana"
	)
	fmt.Println(fruit1)
	fmt.Println(fruit2)
}
