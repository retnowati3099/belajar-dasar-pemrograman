package main

import "fmt"

func main() {
	// menggunakan keyword for
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	fmt.Println("================")

	// menggunakan keyword for dengan argumen hanya kondisi - mirip while
	var i int = 0
	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}

	fmt.Println("================")

	// menggunakan keyword for tanpa argumen
	var j int = 0
	for {
		fmt.Println("Angka", j)
		j++

		if j == 5 {
			break
		}
	}

	// mengguankan keyword for - range

	// keyword break dan continue
}

/*
Di Golang, keyword perulangan hanya for saja, tetapi meskipun demikian kemampuannya merupakan gabungan for, foreach, dan while ibarat bahasa pemrograman lain
*/
