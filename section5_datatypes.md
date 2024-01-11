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
คอนดิชั่นจะทำงานใน () ก่อนเสมอ

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

string คือ []byte{}
thai := "สวัสดี"
ตัวอย่าง y := []byte{0xe0, 0xb8, 0xaa, 0xe0, 0xb8, 0xa7, 0xe0, 0xb8, 0xb1, 0xe0, 0xb8, 0xaa, 0xe0, 0xb8, 0x94, 0xe0, 0xb8, 0xb5}

การหาค่า array string
for i := 0; i < len(thai); i++ {
fmt.Printf(" 0x%x ,", thai[i])
}
