# GO Datatypes

## integers int8 int16 int32

signed 8 16 32 64

unsigned 8 16 32 64 ไม่มีค่าติดลบ

### Two's (2s) Complement เลขฐาน 2

func main() {
x := 127
fmt.Println(bitInt(int8(x)))

}
func bitInt(x int8) [8]byte {
var result [8]byte
for i := 0; i < len(result); i++ {
result[7-i] = byte(x & 1)
x = x >> 1
}
return result
}

## Operators

เรียงอันดับความสำคัญ เรียงตาม president 5 lv 54321 ตามภาพ
คอนดิชั่นจะทำงานในวงเล็บ () จะทำงานก่อนเสมอ

การ mod % เครื่องหมายของผลลัพ 5 , -5 ขึ้นอยู่กับตัวตั้ง

## Floating-point

IEEE 754 Format

## Complex Number

real and imaginary
3 + 4i

ทฤษฎีบทพีทาโกรัส
x ยกกำลัง 2 + y ยกกำลัง 2
ผลลัพ สแครูด
ตัวอย่าง 3*3 9 + 4*4 16 = 25
รูดของ 25 = 5

fmt.Println(cmplx.Abs(complex(3, 4)))

## String

string ในภาษา go คือ []byte{} หรือ array of byte

ascii คือเลข ฐาน 2
dicimal คือเลขฐาน 10
hexdicimal เลขฐาน 16 ขึ้นต้นด้วย 0x

thai := "สวัสดี"
ตัวอย่าง y := []byte{0xe0, 0xb8, 0xaa, 0xe0, 0xb8, 0xa7, 0xe0, 0xb8, 0xb1, 0xe0, 0xb8, 0xaa, 0xe0, 0xb8, 0x94, 0xe0, 0xb8, 0xb5}

การหาค่า array string
for i := 0; i < len(thai); i++ {
fmt.Printf(" 0x%x ,", thai[i])
}

## unicode utf-8

นำตัวอักษรไปค้น unicodelookup.com
นำมาเทียบแปลงค่า utf-8
en.wikipedia.org/wiki/UTF-8
แปลงเลขตามขนาด Byte 1,2, 3 หรือ 4

นำเลข binary ไปแปลงกลับเป็น hexadicimal
hexdicimal เลขฐาน 16 ขึ้นต้นด้วย 0x

fmt.Println("\xe0\xb8\xaa")

## rune index of string

yy := "สวัสดี"
yc := []rune(yy)
fmt.Println(len(yc))
fmt.Printf("%x \n", yc[1]) // เลขฐาน 16
fmt.Println(rune(yc[1])) // เลขฐาน 10
fmt.Printf("%q \n", yc[0]) // แสดงตัวอักษรใน display

### integer to string 6 cases

#### case 1 integer to string

ex1 := string(rune(0x265e))
fmt.Println(ex1)
for i := 0; i < len(ex1); i++ {
fmt.Printf(`\x`+"%x", ex1[i])
}

fmt.Println("\xe2\x99\x9e")

#### case 2 []byte{} to string

ex2 := []byte{0xe2, 0x99, 0x9e}
fmt.Println(string(ex2))

#### case 3 []rune to string

s := "สวัสดี"
// extract rune
ss := []rune(s)
for i := 0; i < len(ss); i++ {
fmt.Printf("0x%x,", ss[i])
}

// []rune to string
sss := []rune{0xe2a, 0xe27, 0xe31, 0xe2a, 0xe14, 0xe35}
fmt.Println(string(sss))

#### case4 slice of byte to string

s := "สวัสดี"
ss := []byte(s)
for i := 0; i < len(ss); i++ {
fmt.Printf(`\x`+"%x", ss[i])
}
fmt.Println()
fmt.Println(string("\xe0\xb8\xaa\xe0\xb8\xa7\xe0\xb8\xb1\xe0\xb8\xaa\xe0\xb8\x94\xe0\xb8\xb5"))

#### case 5 string to slice of rune

x := "วรรณสิงห์"
xx := []rune(x)

for i := 0; i < len(xx); i++ {
fmt.Printf(`0x`+"%x,", xx[i])
}
// slice rune to string
xxx := []rune{0xe27, 0xe23, 0xe23, 0xe13, 0xe2a, 0xe34, 0xe07, 0xe2b, 0xe4c}
fmt.Println(string(xxx))

### Work with string

    x := "ติดสอบ"

xx := len(x)
fmt.Println(xx)

for index, value := range x {
// fmt.Println(string(value))
fmt.Printf("%d %c\n", index, value)

}

#### utf-8

countUtf := utf8.RuneCountInString(x)
fmt.Println("count string :", countUtf)

#### utf-8 decode rune

for i := 0; i < len(x); {
rune, size := utf8.DecodeRuneInString(x[i:])
i = i + size
fmt.Printf("%c-", rune)
}
fmt.Println()

#### ความแตกต่าง bytes package and strings package

//bytes package
fider := "สอ"
fmt.Println(bytes.Count([]byte(x), []byte(fider)))
// strings package
fmt.Println(strings.Count(x, fider))

### String buffer builder

// bytes buffer builder
buff := bytes.Buffer{}
buff.WriteRune('y')
buff.WriteRune('o')
fmt.Println(buff.String())

// string buffer builder
stringBuff := strings.Builder{}
stringBuff.WriteRune('h')
stringBuff.WriteRune('e')
fmt.Println(stringBuff.String())

### string convert

import package : strconv

//ascii to integer
atoi, \_ := strconv.Atoi("123")
fmt.Println(atoi, reflect.TypeOf(atoi))

// integer to acsii
itoA := strconv.Itoa(1234)
fmt.Println(itoA, reflect.TypeOf(itoA))

// string to float64
flost, \_ := strconv.ParseFloat("12.4", 64)
fmt.Println(flost, reflect.TypeOf(flost))

// string to boolean
convertBool, \_ := strconv.ParseBool("t")
fmt.Println(convertBool)

### unicode packages

fmt.Println("digit? :", unicode.IsDigit('1'))
fmt.Println("number? :", unicode.IsNumber(1))
fmt.Println("IsUpper? :", unicode.IsUpper('A'))
fmt.Println("IsLower? :", unicode.IsLower('a'))

## constants

- compiler knows it values
- Evaluation at compile time
  การประกาศตัวแปรรูปแบบ const ระบบจะทำการ Evaluation at times
- Declare constants as a group
- lota
- Untyped constants
