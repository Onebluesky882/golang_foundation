package main

import (
	"fmt"
)

func main() {

	thai := "สวัสดี"
	y := []byte{0xe0, 0xb8, 0xaa, 0xe0, 0xb8, 0xa7, 0xe0, 0xb8, 0xb1, 0xe0, 0xb8, 0xaa, 0xe0, 0xb8, 0x94, 0xe0, 0xb8, 0xb5}

	for i := 0; i < len(thai); i++ {
		fmt.Printf(" 0x%x ,", thai[i])
	}
	fmt.Println()
	fmt.Printf("%s \n", y)
}
