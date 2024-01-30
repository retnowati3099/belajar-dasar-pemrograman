package main

import "fmt"

// deklarasi struct emloyee dengan 2 property, yaitu name dan salary
type employee struct {
	name   string
	salary float32
}

func main() {
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
	fmt.Println("Name of employee 1: ", e1.name)
	fmt.Println("Salary of employee 1: ", e1.salary)

	// inisialisasi object struct
	var e2 = employee{}
	e2.name = "Afriska Yusuf Widyamto"
	e2.salary = 10000000
	fmt.Println("Name of employee 2: ", e2.name)
	fmt.Println("Salary of employee 2: ", e2.salary)

	var e3 = employee{"Queena Putri Askawati", 200000000}
	fmt.Println("Name of employee 3: ", e3.name)
	fmt.Println("Salary of employee 3: ", e3.salary)

	var e4 = employee{name: "Aska Yusuf Pratama"}
	fmt.Println("Name of employee 4: ", e4.name)
	fmt.Println("Salary of employee 4: ", e4.salary)

}

/*
- Deklarasi struct menggunakan keyword "type"
- Semua properti variabel objek pada awalnya memiliki zero value sesuai tipe datanya
- Properti variabel objek bisa diakses menggunakan notasi titik
- Cara inisialisasi variabel object adalah dengan menambahkan kurung kurawal setelah nama struct.
- Nilai masing - masing properti bisa diisi pada saat isialisasi
*/
