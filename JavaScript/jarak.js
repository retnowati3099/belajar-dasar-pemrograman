const prompt = require("prompt-sync")({ sigint: true });

let jarak, kecepatan, waktu;

console.log("Program Menghitung Jarak Berdasarkan Kecepatan dan Waktu");
kecepatan = parseFloat(prompt("Masukkan kecepatan: "));
waktu = parseFloat(prompt("Masukkan waktu: "));

jarak = kecepatan * waktu;

console.log(
	`Jika kecepatan ${kecepatan} km/jam dan waktu tempuh ${waktu} jam, maka jarak yang ditempuh adalah ${jarak} km.`
);
// console.log(
// 	"Jika kecepatan " +
// 		kecepatan +
// 		" km/jam dan waktu tempuh " +
// 		waktu +
// 		" jam, maka jarak yang ditempuh adalah " +
// 		jarak +
// 		" km."
// );
