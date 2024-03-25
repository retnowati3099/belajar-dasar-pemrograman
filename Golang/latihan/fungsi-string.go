package main

import (
	"fmt"
	"strings"
)

func main() {
	// fungsi strings.Contains()
	var isExists = strings.Contains("Retno Wati", "Wati")
	fmt.Println(isExists)
	apakahAda := strings.Contains("Queena Putri Askawati", "Askar")
	fmt.Println(apakahAda)

	// fungsi str]ings.HasPrefix()
	var isPrefix1 = strings.HasPrefix("Retno Wati", "Re")
	fmt.Println(isPrefix1)
	isPrefix2 := strings.HasPrefix("Retno Wati", "Ra")
	fmt.Println(isPrefix2)

	// fungsi strings.ToUpper()
	var str1 = strings.ToUpper("retno")
	fmt.Println(str1)
	str2 := strings.ToUpper("ReTnO")
	fmt.Println(str2)

	// fungsi strings.ToLower()
	var str3 = strings.ToLower("RETNO")
	fmt.Println(str3)
	str4 := strings.ToLower("RetNO")
	fmt.Println(str4)

	// fungsi strings.HasSuffix()
	var isSuffix1 = strings.HasSuffix("Retno", "no")
	fmt.Println(isSuffix1)
	isSuffix2 := strings.HasSuffix("Retno", "na")
	fmt.Println(isSuffix2)

	// fungsi strings.Count()
	var howMany1 = strings.Count("Retno Wati", "t")
	fmt.Println(howMany1)
	howMany2 := strings.Count("Afriska Yusuf Widyamto", "u")
	fmt.Println(howMany2)

	// fungsi strings.Index()
	var index1 = strings.Index("Retno", "n")
	fmt.Println(index1)
	index2 := strings.Index("Retno Wati", "t")
	fmt.Println(index2)

	// fungsi strings.Replace()
	var text = "banana"
	var find = "a"
	var replaceWith = "o"
	fmt.Println(strings.Replace(text, find, replaceWith, 1))
	fmt.Println(strings.Replace(text, find, replaceWith, 2))
	fmt.Println(strings.Replace(text, find, replaceWith, -1))

	// fungsi strings.Repeat()
	var str5 = strings.Repeat("na", 3)
	fmt.Println(str5)
	str6 := strings.Repeat("na", 5)
	fmt.Println(str6)

	// fungsi strings.Split()
	var string1 = strings.Split("The dark night", " ")
	fmt.Println(string1)
	string2 := strings.Split("banana", "")
	fmt.Println(string2)

	// fungsi strings.Join()
	var data1 = []string{"banana", "papaya", "tomato"}
	var string3 = strings.Join(data1, "-")
	fmt.Println(string3)
	data2 := []string{
		"banana",
		"papaya",
		"tomato",
	}
	string4 := strings.Join(data2, "_")
	fmt.Println(string4)
}

/*
	- Fungsi string.Contains() digunakan untuk mendeteksi apakah string (parameter keedua) merupakan bagian dari string lain(parameter pertama). Nilai kembaliannya berupa bool.
	- Fungsi str]ings.HasPrefix() digunakan untuk mendeteksi apakah sebuah string (parameter pertama) diawali string tertentu (parameter kedua.
	- Fungsi strings.ToUpper() digunakan untuk mengubah huruf - huruf string menjadi huruf besar.
	- Fungsi strings.ToLower() digunakan untuk megubah huruf - huruf string menjadi huruf - huruf kecil.
	- Fungsi strings.HasSuffix() digunakan untuk mendeteksi apakah sebuah string(parameter pertama) diakhiri string tertentu (parameter kedua)
	- Fungsi strings.Count() digunakan untuk menghitung jumlah karakter tertentu (parameter kedua) dari sebuah string (parameter pertama). Nilai kembaliannya adalah jumlah karakter.
	- Fungsi strings.Index() digunakan untuk mencari posisi indeks dari sebuah string (parameter kedua) dalam string (parameter pertama). Jika ditemukan 2 substring atau lebih, maka yang diambil adalah yang pertama.
	- Fungsi strings.Replace() digunakan untuk replace atau mengganti bagian dari string dengan string tertentu. Jumlah substr]ingyang di-replace bisa ditentukan, apakah hanya 1 sting pertama, 2 string atau seluruhnya. Jika ingin replace pada semua substr]ing, maka parameter keempat diisi angka -1.
	- Fungsi strings.Repeate() digunakan untuk mengulang string (parameter pertama) sebanyak data yang ditentukan (parameter kedua)
	- Fungsi str]ings.Split() digunakan untuk memisahkan string (parameter pertama) dengan tanda pemisah bisa ditentukan sendiri (parameter kedua). Hasilnya berupa slice string.
	- Fungsi strings.Join() memiliki kegunaan berkebalikan dengan strings.Split(). Digunakan untuk menggabungkan slice string (parameter pertaa) menjadi sebuah str]ing dengan pemisah tertentu (parameter kedua)
*/
