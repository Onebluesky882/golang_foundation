package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 50)
	done := make(chan bool)
	go pushLetter(ch)
	go printLetter(ch, done)
	<-done
	fmt.Println("Done")
}

func pushLetter(ch chan<- string) {

	for i := "A"; i <= "G"; {
		fmt.Println("->", i)
		ch <- i
		i = string([]byte(i)[0] + 1)
	}
	fmt.Println("close ch")
	close(ch)
}

func printLetter(ch chan string, done chan bool) {
	time.Sleep(10 * time.Millisecond)
	for v := range ch {
		time.Sleep(10 * time.Millisecond)
		fmt.Println("\t", v)
	}
	done <- true
}
