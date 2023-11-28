const prompt = require("prompt-sync")({ sigint: true });
let nilai = prompt("Masukkan nilai kamu: ");

if (nilai <= 100 && nilai >= 85) {
	console.log(`nilai kamu ${nilai}, kamu mendapat A.`);
} else if (nilai >= 70) {
	console.log(`nilai kamu ${nilai}, kamu mendapat B.`);
} else if (nilai >= 40) {
	console.log(`nilai kamu ${nilai}, kamu mendapat C.`);
} else if (nilai >= 20) {
	console.log(`nilai kamu ${nilai}, kamu mendapat D.`);
} else {
	console.log(`nilai kamu ${nilai}, kamu mendapat E.`);
}
