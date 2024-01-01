const prompt = require("prompt-sync")({ sigint: true });
let a = parseInt(prompt("a = "));
let b = parseInt(prompt("b = "));

// function declaration
console.log(solveMeFirst(a, b));
function solveMeFirst(a, b) {
	return a + b;
}
console.log(solveMeFirst(a, b));

// function ...
// const solveMeFirst = function (a, b) {
// 	return a + b;
// };
// console.log(solveMeFirst(a, b));

// arrow function
// const solveMeFirst = () => {
// 	return a + b;
// };
// console.log(solveMeFirst(a, b));
