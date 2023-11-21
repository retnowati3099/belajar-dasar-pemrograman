const prompt = require("prompt-sync")({ sigint: true });
let num1 = parseInt(prompt("Masukkan bilangan pertama: "));
let num2 = parseInt(prompt("Masukkan bilangan kedua: "));
let jumlah = num1 + num2;
console.log(num1 + " + " + num2 + " = " + jumlah);
