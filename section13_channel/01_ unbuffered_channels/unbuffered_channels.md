# 01_unbuffered_channels

(synchronous channels)

```go

package main
import (
 "fmt"
)
func main() {
 ch := make(chan int)
 done := make(chan string) // done <- struct{}{}
 chSquare := make(chan int)

 go sander(ch, done)
 go receiver(chSquare, done)
 go square(ch, chSquare, done)

 for v := range done {
  fmt.Println("Done status : ", v)
 }

 fmt.Println("Main exit")

}

// sander
func sander(ch chan int, done chan string) {
 for i := 0; i <= 5; i++ {
  fmt.Println("sending value", i)
  ch <- i
 }
 close(ch)
 fmt.Println("Sander is  about to complete")
 done <- "done from sander"
 fmt.Println("Sander done")

}

//square
func square(ch chan int, chSquare chan int, done chan string) {

 for v := range ch {

  chSquare <- v * v
 }
 close(chSquare)
 done <- "done from square"
}


// receiver
func receiver(ch chan int, done chan string) {

 for v := range ch {
  fmt.Println("\t Receive Value", v)
 }

 done <- "done from reciever"
 close(done)
}

```
