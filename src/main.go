package main

import "fmt"

func printEachLine(args ...interface{}) {
	for _, v := range args {
		fmt.Println(v)
	}
}

func main() {
	x := []interface{}{"asdfwf", "wfwfd", 233, "wfwf"}
	printEachLine(x)
	printEachLine(x...)

}
