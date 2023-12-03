const prompt = require("prompt-sync")({ sigint: true });

let panjang = parseInt(prompt("Masukkan panjang tabel perkalian: "));
let lebar = parseInt(prompt("Masukkan lebar tabel perkalian: "));

process.stdout.write("*\t");
for (let i = 1; i <= panjang; i++) {
	process.stdout.write(i + "\t");
}
console.log();
for (let i = 1; i <= lebar; i++) {
	process.stdout.write(i + "\t");
	for (let j = 1; j <= panjang; j++) {
		process.stdout.write(i * j + "\t");
	}
	console.log();
}
