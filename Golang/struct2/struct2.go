package main

import "fmt"

type fruit struct {
	color string
}

type human struct {
	name string
	age  int
}

// struct human diembed ke dalam struct student, caranya yaitu dengan menuliskan nama struct yang ingin diembed ke dalam body struct target
type student struct {
	human
	age   int
	grade float32
}

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

	// variabel object pointer
	var e5 = employee{name: "cantika putri askawati", salary: 8000000}
	var e6 *employee = &e5 // e6 adalah variabel pointer hasil cetakan struct employee. e6 menampung nilai referensi e5
	fmt.Println(e5)
	fmt.Println(e6)

	//embeded struct
	var s1 = student{}
	s1.name = "Alika Prettia Askawati"
	s1.age = 21 // age of student
	s1.grade = 4.00
	s1.human.age = 20 // age of human

	fmt.Println("Student name: ", s1.name)
	fmt.Println("Student name: ", s1.human.name)
	fmt.Println("Student age: ", s1.age)
	fmt.Println("Student age: ", s1.human.age)
	fmt.Println("Student grade: ", s1.grade)

	var h2 = human{
		name: "Afifah",
		age:  6,
	}
	var s2 = student{human: h2, age: 20, grade: 4.00}
	fmt.Println("Name: ", s2.name)

	// anonymous struct tanpa pengisian properti
	var pinapple = struct {
		fruit
		num int
	}{}

	pinapple.fruit = fruit{"yellow"}
	pinapple.num = 10

	fmt.Println("color: ", pinapple.fruit.color)
	fmt.Println("num: ", pinapple.num)

	// anonymous struct dengan pengisian properti
	var apple = struct {
		fruit
		num int
	}{
		fruit: fruit{color: "red"},
		num:   5,
	}

	fmt.Println("color: ", apple.fruit.color)
	fmt.Println("num: ", apple.num)

}

/*
- Deklarasi struct menggunakan keyword "type"
- Semua properti variabel objek pada awalnya memiliki zero value sesuai tipe datanya
- Properti variabel objek bisa diakses menggunakan notasi titik
- Cara inisialisasi variabel object adalah dengan menambahkan kurung kurawal setelah nama struct.
- Nilai masing - masing properti bisa diisi pada saat isialisasi
- Embedded struct adalah mekanisme untuk menempelkan sebuah struct sebagai properti struct lain
- Embedded struct bersifat mutable, nilai propertinya bisa berubah.
- Khusus untuk properti yang bukan properti asli (properti turunan dari struct lain), bisa diakses dengan mengakses struct parent terlebih dahulu.
- Jika salah satu nama properti sebuah struct memiliki kesamaan dengan properti milik struct lain yang di-embed, maka pengaksesan propertinya harus dilakukan secara eksplisit atau jelas
- Pengisian nilai proprty substruct bisa dilakukan dengan langsung memasukkan variabl objek yang tercetak dari struct yang sama
- Anonymous struct adalah struct yang tidak didklarasikan di awal sebagai tipe data baru, melainkan langsung ke pembuatan objek
- Dalam pembuatan anonymous struct adalah deklarasi harus diikuti dengan inisialisasi.
Pada contoh program di atas, pada pineappl stlah deklarasi struktur struct, terdapat kurung kurawal untuk inisialisasi objek. MMeskipun nilai tidak diisikan di awal, kurung kurawal tetap harus ditulis
*/
