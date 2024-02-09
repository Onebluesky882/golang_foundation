# Unidirectional channels type

( Unidirectional channel types)
channel ที่ทำหน้าที่ได้อย่างเดียว ส่งได้อย่างเดียว หรือรับได้อย่างเดียว

Send only

chan<-

func f(ch chan<- int){
...
}

Receive only

<-chan

func f(ch <-chan int){
...
}

## Unidirectional channel type

จะไม่สามารถรับ chanal แบบ synchronous ได้

## sander chan<-

ขาส่งเท่านั้นจะเป็นตัวปิด channel
close(ch)

receiver <-chan ขารับไม่สามารถทำการปิดแชแนลได้

```go
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

```
