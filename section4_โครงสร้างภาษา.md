# Golearn foundation

## fmt.Printf

%#v = value
%T = type

## pointer

declear var p \*int

p = 40
fmt.Println(&p) for watch address
fmt.Println(\*p) for watch value

## func return pointer

func fb() \*int {
x := 4
return &x
}

func main() {
c := fb()
fmt.Println(\*c)
}

## function receive pointer คอนเซป ของ alias

alias คือ นามแฝง
func fb(x *int) {
*x = _x_ 2
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

### การ comparison

การเปรียบเทียบตัวแปร ต้องมี type เดียวกัน
KG == KG
LB == LB

#### การ class ตัวแปรให้เปลี่ยน type

func main() {
k := KG(3)
b := LB(3)
c := KG(b)
fmt.Println(k == c)
}

#### การแปลง type ด้วย func ให้สามารถคอมแพกันได้ associate

type KG float64
type LB float64
type ST string

func (b LB) toKG() KG {
return KG(b / 2.20462)
}

func (kg KG) kgToPound() LB {
return LB(kg / 0.453592)
}

func (kg KG) toString() ST {
return ST(strconv.Itoa(int(kg)) + " kg")
}

### if else

Logical operators เรียงจากลำดับความสำคัญใหญ่ลงเล็ก
&&
||
!

### control flow switch case

switch (condition) {
case 1 :
case 2 :
}

### for loop

// for initialization; condition ; post{
body
}

//for (condition) {
body
}
