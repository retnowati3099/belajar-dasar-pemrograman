const prompt = require("prompt-sync")({ sigint: true });

let baris = parseInt(prompt("Masukkan banyak baris: "));

for (let i = 1; i <= baris; i++) {
	for (let j = 1; j <= i; j++) {
		process.stdout.write("* ");
	}
	console.log();
}

console.log("===========");

let bintang = "";
for (let i = 1; i <= baris; i++) {
	bintang += "* ";
	console.log(bintang);
}

console.log("===========");

for (let i = 1; i <= baris; i++) {
	// spasi
	for (let j = baris - 1; j >= i; j--) {
		process.stdout.write(" ");
	}

	// bintang
	for (let k = 1; k <= i; k++) {
		process.stdout.write("* ");
	}
	console.log();
}
