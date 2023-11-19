const prompt = require("prompt-sync")({ sigint: true });
let bilangan = parseInt(prompt("Masukkan Bilangan: "));

if (bilangan % 2 == 0) {
  console.log(bilangan + " merupakan bilangan genap.");
  //   console.log(`${bilangan} merupakan bilangan genap.`);
} else {
  console.log(bilangan + " merupakan bilangan ganjil.");
  //   console.log(`${bilangan} merupakan bilangan ganjil.`);
}
