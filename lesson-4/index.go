package main

import "fmt"

// Go usa o padrão IEEE 754 para pontos flutuantes, não é recomendado usaló para calculos monetários nem mesmo comparação.
//
//	Isso ocorre por quê alguns valores em decimal não possuem representação exata em binário.

func main() {
	var smallFloat float32
	smallFloat = 2.1522
	fmt.Println(smallFloat)

	var bigFloat float64
	bigFloat = 43.123456789321654
	fmt.Println(bigFloat)

	var myComplex complex128
	myComplex = complex(bigFloat, bigFloat)
	fmt.Println(myComplex)
}
