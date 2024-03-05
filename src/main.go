package main

import "fmt"

func createFib() func(int) []int {
	fList := []int{0, 1, 1}
	return func(n int) []int {
		if n > len(fList) {
			for n > len(fList) {
				lastIndex := len(fList) - 1
				fList = append(fList, fList[lastIndex]+fList[lastIndex-1])
			}
		}
		return fList[:]
	}

}

func main() {

	fib := createFib()

	fmt.Println(fib(5)) // 0 1 1 2 3
	fmt.Println(fib(6)) // 0 1 1 2 3 5
	fmt.Println(fib(7)) // 0 1 1 2 3 5 8

	fmt.Println(fib(16))

}
