const PHI = 3.14;
const prompt = require("prompt-sync")({ sigint: true });
let radius = parseFloat(prompt("Masukkan radius: "));
let luas = PHI * radius * radius;
console.log(
  `Lingkaran yang memiliki radius sebesar ${radius} memiliki luas ${luas}`
);
