package main

import "fmt"

// strings
func main() {
	var myString string
	myString = "Hello\n2026"
	fmt.Println(myString)

	// se usar `` ele imprime a tabulação setada
	myString = `Hello 
        
	2026`

	fmt.Println(myString)

	var fistName, lastName string
	fistName = "Vitor Manoel"
	lastName = "Porfirio"
	fmt.Println("My name is = " + fistName + " " + lastName)
	//ou
	fmt.Printf("%s %s\n", fistName, lastName)

	var fullName string
	fullName = fistName + " " + lastName
	fmt.Println(fullName)

}
