package main

import "fmt"

func main() {
	ch := make(chan int)
	done := make(chan string)

	go sender(ch)
	go receiver(ch, done)
	<-done
	fmt.Println("done")
}

func sender(ch chan<- int) {
	ch <- 88

	close(ch)
}

func receiver(ch <-chan int, done chan string) {
	fmt.Println("Receive", <-ch)
	done <- "true"
	close(done)
}
