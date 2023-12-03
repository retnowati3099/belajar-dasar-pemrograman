const prompt = require("prompt-sync")({ sigint: true });

let baris = parseInt(prompt("Masukkan banyak baris: "));

// for (let i = 1; i <= baris; i++){
//     for (let j = 1; j <= i; j++) {
//         console.log("*");
//     }
// }

let bintang = "";
for (let i = 1; i <= baris; i++) {
	bintang += "* ";
	console.log(bintang);
}

console.log("===========");
