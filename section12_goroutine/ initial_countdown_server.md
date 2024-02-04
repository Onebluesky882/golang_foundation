# initial countdown server

```go

package main

import (
 "io"
 "log"
 "net"
 "time"
)

func main() {
 log.SetFlags(log.Ltime)
 listener, err := net.Listen("tcp", "localhost:8080")
 if err != nil {
  log.Fatal(err)
 }
 for {
  conn, err := listener.Accept()
  if err != nil {
   log.Fatal(err)
   return
  }
  // logic
  countingDownHandler(conn)
 }

}

func countingDownHandler(conn net.Conn) {
 defer func() {
  io.WriteString(conn, "your connetion will be close...\n")

  conn.Close()
 }()
 count := 5
 for {
  io.WriteString(conn, " haha123 \n")
  time.Sleep(time.Second)
  count--
  if count == 0 {
   io.WriteString(conn, "Enter your number \n")
   break
  }
 }

}

```
