package main

import "fmt"

// deklarasi struct emloyee dengan 2 property, yaitu name dan salary
type employee struct {
	name string
	salary float32
}

func main(){
	// membuat variabel objek -> sama seperti pembuatan variabel biasa
	// tulis nama variabel diikuti nama struct
	var e1 employee

	// menampilkan isi properti varibel objek yang belum diisi
	fmt.Println(e1.name)
	fmt.Println(e1.salary)

	// mengisi properti variabel objek
	e1.name = "Retno"
	e1.salary = 2400000

	// menampilkan isi properti varibel objek yang sudah diisi
	fmt.Println("Name: ", e1.name)
	fmt.Println("Salary: ", e1.salary)
}

/*
- Deklarasi struct menggunakan keyword "type"
- Semua properti variabel objek pada awalnya memiliki zero value sesuai tipe datanya
- Properti variabel objek bisa diakses menggunakan notasi titik
*/