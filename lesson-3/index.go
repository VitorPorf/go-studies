package main

import "fmt"

// variaveis do tipo unsigned não possuem o bit do sinal, por exemplo = 2^8, se for assigned, seria 2^7 1 bit representa o sinal

func main() {
	var unsignedSmallInt uint8 // max 255 - 8 bits
	unsignedSmallInt = 200
	fmt.Println(unsignedSmallInt)

	var smallAssignedInt int8
	smallAssignedInt = -127
	println(smallAssignedInt)

	var smallNegativePositive int
	smallNegativePositive = int(smallAssignedInt)
	println(smallNegativePositive)
}
