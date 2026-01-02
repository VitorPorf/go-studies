package main

import (
	"fmt"
	"strconv"
)

// Operadores logicos := OR,XOR,AND ...
func main() {
	var a uint8 = 201
	var b uint8 = 105

	c := a | b

	//OR a saida vai ser 233, pq vai ser bit a bit
	//  11001001  (201)
	// 	01101001  (105)
	//-----------
	// 	11101001 (233)

	fmt.Println(c)

	//AND a saida vai ser 73, pq vai ser bit a bit
	//  11001001  (201)
	// 	01101001  (105)
	//-----------
	// 	01001001 (73)
	d := a & b
	fmt.Println(d)

	//XOR a saida vai ser 160, pq vai ser bit a bit
	//  11001001  (201)
	// 	01101001  (105)
	//-----------
	// 	10100000 (160)
	e := a ^ b
	fmt.Println(e)

	//NAND
	//  11001001  (201)
	// 	01101001  (105)    2 + 8 = 10
	//-----------
	// 	10110110 (160) 128 + 32 + 16 + 4 + 2 = 182
	nand := a &^ b
	fmt.Println(nand)

	fmt.Println(strconv.FormatUint(uint64(a), 2))
	fmt.Println(strconv.FormatUint(uint64(^a), 2))

}
