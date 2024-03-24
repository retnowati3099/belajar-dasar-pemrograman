package main

import "fmt"

func main() {
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari:", chicken["januari"])
	fmt.Println("maret:", chicken["maret"])

	// inisialisasi nilai map
	var data map[string]int
	data = map[string]int{}
	data["one"] = 1

	// mendefinisikan nilai map dengan cara horisontal
	var chicken1 = map[string]int{"januari": 50, "februari": 60}
	fmt.Println(chicken1)
	fmt.Println(chicken1["januari"])

	// mendefinisikan nilai map dengan cara vertikal
	var chicken2 = map[string]int{
		"januari":  70,
		"februari": 80,
	}
	fmt.Println(chicken2)
	fmt.Println(chicken2["januari"])

	// inisialisasi nilaim map tanpa nilai awal
	var chicken3 = map[string]int{}
	var chicken4 = make(map[string]int)
	var chicken5 = *new(map[string]int)
	fmt.Println(chicken3)
	fmt.Println(chicken4)
	fmt.Println(chicken5)

	// iterasi item map menggunakan for - range
	var chicken6 = map[string]int{
		"januari":  100,
		"februari": 150,
		"maret":    200,
		"april":    250,
	}

	for key, val := range chicken6 {
		fmt.Println(key, val)
	}

	// panjang map
	fmt.Println(len(chicken6))

	// menghapus item map
	delete(chicken6, "april")
	fmt.Println(len(chicken6))
	fmt.Println(chicken6)

	// deteksi keberadaan item dengan key tertentu
	var value, isExist = chicken6["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("Item is not exist")
	}

	// kombinasi slice dan map
	// slice yang tipe tiap elemen-nya adalah map[string]int
	var chickens = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}
	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

	var pets = []map[string]string{
		{"name": "punpun", "gender": "female"},
		{"name": "pinpin", "gender": "female"},
		{"name": "ponpon", "gender": "male"},
	}
	for _, pet := range pets {
		fmt.Println(pet["name"], pet["gender"])
	}
}

/*
	- Map adalah tipe data asosiatif yang ada di Go, berbentuk pasangan key-value untuk setiap data (atau value) yang disimpan.
	- Key harus unik karena digunakan sebagai penanda (identifier) untuk pengaksesan value yang bersangkutan.
	- Cara menggunakan map cukup dengan menuliskan keyword map diikuti tipe data key dan valuenya.
	- Default nilai variabel map adalah nil, sehingga perlu dilakukan inisialisasi nilai default di awal, caranya dengan menambahkan kurung kurawal pada akhir tipe.
	- Cara menge-setnilai pada sebuah map adalah dengan menuliskan variabelnya, kemudidan disisipkan key pada kurung sikuvariabel, lalu isi nilainya.
	- Pengisian data pada map bersifat overwrite, ketika variabel sudah memiliki item denfan key yang sama, maka value lama akan ditimpa dengan value yang baru.
	- Pada pengaksesan item menggunakan key yang belum tersimpan di map, akan dikembalikan nilai default tipe data valuenya.
	- Niai variabel bertipe map bisa didefinisikan di awal, caranya dengan menambahkan kurung kurawal setelah tipe data, lalu menuliskan key dan value di dalamnya.
	- Key dan value dituliskan dengan pembatas tanda titik dua, sedangkan tiap itemnya dituliskan dengan pembatas koma, khusus deklarasi dengan gaya verikal, tanda koma perlu dituliskan setelah item terakhir.
	- fungsi delete() digunakan untuk menghapus item denfan key tertentu pada variabel map. Cara penggunaannya dengan memasukkan objek map dan key item yang ingin dihapus sebagai parameter.
	- Cara untuk mengetahui apakah dalam sebuah variabel map terdapat item dengan key tertentu atau tidak, yaitu dengan memanfaatkan 2 variabel sebagai penampung nilai kembalian pengaksesan item. Return value ke-2 ini adalah opsional, isinya nilai bool yang menunjukkan ada atau tidaknya item yang dicari.
	- dalam []map[string]string, tiap elemen bisa saja memiliki key yang berbeda - beda.
*/
