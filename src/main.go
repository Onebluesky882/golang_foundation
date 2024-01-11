package main

import (
	"fmt"
)

func main() {
	// case 1 integer to string
	// integer -> string
	ex1 := string(rune(0x265e))
	y := "\xe2\x99\x9e"
	for i := 0; i < len(ex1); i++ {
		fmt.Printf(`\`+"x%x", ex1[i])
	}
	fmt.Println()
	fmt.Println(y == ex1)

	//case2 slice of byte to string
	//[]byte -> string
	ex2 := []byte{0xe2, 0x99, 0x9e}
	ex2String := string(ex2)
	fmt.Println(ex2String)

	//case3 slice of rune to string
	// []rune -> string
	// ex3 := []rune{0xe0, 0xb8, 0xaa, 0xe0, 0xb8, 0xa7, 0xe0, 0xb8, 0xb1, 0xe0, 0xb8, 0xaa, 0xe0, 0xb8, 0x94, 0xe0, 0xb8, 0xb5}
	yy := "สวัสดี"
	yc := []rune(yy)
	fmt.Println(len(yc))
	fmt.Printf("%x \n", yc[1])
	fmt.Printf("%q \n", yc[0])
	fmt.Println(rune(yc[1]))
}
