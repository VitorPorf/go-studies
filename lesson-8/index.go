package main

import (
	"fmt"
	"strconv"
)

// Bitwise Shift
// 0101 | (1<<1) resultado em =
// 0101
// (0001, representa 1 em binário, move o bit a esquerda 0010)
// 0101
// 0010
// ------
// 0111
func main() {
	var a uint8 = 20
	fmt.Println(strconv.FormatUint(uint64(a), 2))
	result := a >> 2
	fmt.Println(strconv.FormatUint(uint64(result), 2))

	result = a << 2
	fmt.Println(strconv.FormatUint(uint64(result), 2))
}
