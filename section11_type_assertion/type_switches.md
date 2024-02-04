# type switches

```go

package main

import (
 "fmt"
)

func switchtestValue(x interface{}) {
 switch v := x.(type) {
 case string:
  fmt.Println("string value", v)
 case int:
  fmt.Println("int value", v)
 case bool:
  fmt.Println("bool value", v)
 default:
  fmt.Println("not match any type")
 }
}

func iftestValue(x interface{}) {
 if v, ok := x.(string); ok {
  fmt.Println("string value", v)
  return
 }
 if v, ok := x.(int); ok {
  fmt.Println("int value", v)
  return
 }
 if v, ok := x.(bool); ok {
  fmt.Println("bool value", v)
  return
 }
 if v, ok := x.(float32); ok {
  fmt.Println("float32 value", v)
  return
 }
 if v, ok := x.(float64); ok {
  fmt.Println("float64 value", v)
  return
 }
 fmt.Println(" not match any match")
}

func main() {
 switchtestValue("hello")
 switchtestValue(3)
 switchtestValue(true)
 switchtestValue(float32(3.3))
 switchtestValue(3.11)
}

```
