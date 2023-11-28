const prompt = require("prompt-sync")({ sigint: true });

let bil1, bil2, bil3;

bil1 = parseInt(prompt("Masukkan bilangan 1: "));
bil2 = parseInt(prompt("Masukkan bilangan 2: "));
bil3 = parseInt(prompt("Masukkan bilangan 3: "));

if (bil1 >= bil2 && bil1 >= bil2) {
	console.log(`Bilangan terbesar adalah ${bil1}`);
} else if (bil2 >= bil1 && bil2 >= bil3) {
	console.log(`Bilangan terbesar adalah ${bil12}`);
} else {
	console.log(`Bilangan terbesar adalah ${bil3}`);
}
