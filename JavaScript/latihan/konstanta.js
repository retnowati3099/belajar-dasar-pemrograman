const PI = 3.14;
console.log("Phi: ", PI);
{
	const PI = 22 / 7;
	console.log("Phi: ", PI);
}

// membuat array konstanta
const fruits = ["apple", "pineapple", "banana"];

// mengubah elemen
fruits[2] = "papaya";

// menambahkan elemen
fruits.push("watermelon");

// membuat objek konstanta
const student = {
	name: "Retno",
	age: 23,
	grade: 3.53,
};

// mengubah properti
student.name = "Retno Wati";

// menambahkan properti
student.gender = "Female";

/*
    - Variabel yang didefinisikan dengan const tidak dapat dideklarasi ulang, tidak dapat diassign ulang, dan memiliki cakupan blok.
    - Gunakan const saat mendeklarasikan array baru, objek baru, fungsi baru, dan regexp baru.
    - Tidak bisa menetapkan ulang array maupun objek
*/
