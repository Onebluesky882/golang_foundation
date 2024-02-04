# type assertion

```go
// การ assertion type สังเกตุ w.(*os.File)
package main

import (
 "fmt"
 "io"
 "os"
)

func main() {
 var w io.Writer
 w = os.Stdout
 result, ok := w.(*os.File)

 fmt.Printf("%T %#v\n", w, w)
 fmt.Printf("%T %#v %v\n", result, result, ok)
}

```
