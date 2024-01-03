# Golearn foundation

## fmt.Printf

%#v = value
%T = type

## pointer

declear var p *int

p = 40
fmt.Println(&p) for watch address
fmt.Println(*p) for watch value

## func return pointer

func fb() *int {
x := 4
return &x
}

func main() {
c := fb()
fmt.Println(*c)
}

## function receive pointer คอนเซป ของ alias

alias คือ นามแฝง
func fb(x *int) {
*x =  *x* 2
}

func main() {
y := 2
fb(&y)
fmt.Println(y)
}

### (new function)

 a := new(int)
 fmt.Printf("%T", a)

## type declaration

type YourTypeName float64
