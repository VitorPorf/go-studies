package main

import "fmt"

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
