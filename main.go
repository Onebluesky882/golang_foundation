package main

import "fmt"

type KG float64
type LB float64

func main() {
	k := KG(3)
	l := LB(3)
	fmt.Println(k == 3)
	fmt.Println(l == 3)
}
