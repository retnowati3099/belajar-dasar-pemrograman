package main

import "fmt"

// deklarasi banyak constant di luar fungsi main
const (
	FIRSTNAME string = "Retno"
	LASTNAME         = "Wati"
)

func main() {
	const PI = 3.14
	fmt.Println("Nilai phi adalah ", PI)

	fmt.Println("First name: ", FIRSTNAME)
	fmt.Println("Last name: ", LASTNAME)

	// deklarasi banyak constant di dalam fungsi main
	const (
		A string = "Retno"
		B        = "Wati"
	)
	fmt.Println("A: ", A)
	fmt.Println("B: ", B)
}
