const prompt = require("prompt-sync")({ signint: true });

let panjang;
let sum = 0;

panjang = parseInt(prompt("Masukkan jumlah angka yang ingin dijumlahkan: "));

let temp = [panjang];

for (let i = 0; i < panjang; i++) {
	temp[i] = parseInt(prompt("Angka ke-" + (i + 1) + ": "));
	sum += temp[i];
}
console.log(sum);
