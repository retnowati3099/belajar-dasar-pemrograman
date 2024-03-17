// deklarasi variabel secara langsung
fullName = "Retno Wati";

// deklarasi variabel menggunakan keyword var
var firstName = "Retno";
var lastName;
lastName = "Wati";

console.log("Hello " + firstName + " " + lastName + "!");
console.log(`Hello ${firstName} ${lastName}!`);
console.log("Hello %s %s!", firstName, lastName);
process.stdout.write("Hello " + firstName + " " + lastName + "!\n");
console.log("Hello", firstName, lastName, "!");

// deklarasi variabel menggunakan keyword let
let student1 = "Afriska";
let student2;
student2 = "Retno";

// deklarasi banyak variabel
let first = "satu",
	second = "dua",
	third = "tiga";
console.log(first, second, third);

let fourth, fifth, sixth;
fourth = "empat";
fifth = "lima";
sixth = "enam";
console.log(fourth, fifth, sixth);

/*
    var -> scopenya funcional atau global
    let -> scopenya block
*/
