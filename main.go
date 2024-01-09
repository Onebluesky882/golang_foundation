package main

import (
	"fmt"
	"strconv"
)

type KG float64
type LB float64
type ST string

func main() {
	k := KG(5.9)
	b := LB(2.20462)

	fmt.Println(k == b.toKG())
	fmt.Println(k.toLB())
	fmt.Println(k.toString())
}

func (lb LB) toKG() KG {

	return KG(lb / 2.20462)
}

func (kg KG) toLB() LB {
	return LB(kg / 0.453592)
}

func (kg KG) toString() ST {
	return ST(strconv.Itoa(int(kg)) + " kg")
}
