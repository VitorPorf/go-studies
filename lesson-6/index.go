package main

import "fmt"

var global = 10

func printGlobal() {
	fmt.Println(global)
}

// Declaração dde variaveis
func main() {
	age := 23
	size, name := 1.77, "Vitor Manoel"
	var lang, address = "Portugues\n", "Goiania, jardim america"

	fmt.Println(age, size, name)
	fmt.Println(lang, address)
	printGlobal()
}
