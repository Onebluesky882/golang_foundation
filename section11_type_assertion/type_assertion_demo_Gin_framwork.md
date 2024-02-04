# gin demo type assertion

```go
package main

import "fmt"

type MyIo struct{}

func Value(key interface{}) {

 if keyAsString, ok := key.(string); ok {
  fmt.Println("assert ok value is ", keyAsString)
  return
 }
 fmt.Println("assert not ok")
}

func main() {
 Value("my name")
}
```
