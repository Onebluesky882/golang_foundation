package main

import (
	"fmt"
)

func main() {
	// map คือ key คือ hash table ของ value
	//### map create

	// แบบที่ 1 using map literal styles
	item := map[string]int{
		"pen":   4,
		"pecil": 5,
		"a":     0,
		"b":     1,
		"c":     2,
	}

	// แบบที่ 2 using make function styles
	orders := make(map[string]int)
	orders["pen"] = 16
	orders["pencil"] = 20

	//### Accessing map's value
	fmt.Println("pen :", item["pen"])
	fmt.Println(orders)

	//### การ delete object in map
	// delete(item, "pen")

	//### add object in map
	x := item
	x["ruler"] = 5
	fmt.Println(x)
	fmt.Println(item)

	// ### find item in map
	value, ok := item["pen"]
	if !ok {
		fmt.Println("item not exist")
	} else {
		fmt.Println(value)
	}

	// การ map จะไม่เรียง ตาม orders
	for key, value := range item {
		fmt.Println(key, value)
	}
}
