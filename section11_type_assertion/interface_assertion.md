# interface assertion

ค่าที่ return กลับมาเป็น interface



```go
package main

import (
 "fmt"
 "io"
)

type Myto struct{}

func (m Myto) Write(p []byte) (n int, err error) {
 return len(p), nil
}

func main() {
 var w io.Writer
 w = Myto{}

 //
 result, ok := w.(io.Writer)

 fmt.Printf("%T %#v\n", w, w)
 fmt.Printf("%T %#v %v\n", result, result, ok)
}


```
