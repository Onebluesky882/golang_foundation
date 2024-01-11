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
