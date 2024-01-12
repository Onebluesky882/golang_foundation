package main

import (
	"fmt"
)

type Peice int

func main() {
	const (
		Sunday = iota + 1 // constant generator
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday, Monday, Tuesday)
}
