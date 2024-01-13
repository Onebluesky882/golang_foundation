package main

import "fmt"

//Type struct
type Person struct {
	name string
	age  int
}

func main() {
	// แบบที่ 1
	staff := Person{name: "wansing", age: 20}

	fmt.Println(staff)

}
