package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

// function dengan parameter
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Helo", firstName, lastName)
}

// function return value
func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

// return multiple value
func getFullName() (string, string){
	return "Retno", "Wati"
}

func main() {
	sayHello()
	sayHelloTo("Retno", "Wati")

	result := getHello("Aska")
	fmt.Println(result)
	fmt.Println(getHello("Aska"))

	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	// menghiraukan return value
	namaDepan, _ := getFullName()
	fmt.Println(namaDepan)
}
