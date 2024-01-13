# Data Structure

## Array

    fruit := [5]string{"apple", "banana", "papaya", "orange", "mango"}
    fmt.Println(fruit[0])
    // type array 2 slots != type array 3 slots
    // ตัวอย่าง
    twoSlots := [3]int{1, 2, 3}
    threeSlots := [3]int{1, 2, 3}
    fmt.Println(twoSlots == threeSlots)

    fruits := [...]string{
    "papaya",
    "banana",
    "mango",
    "apple",
    "orange",
    }
    fmt.Println(fruits, reflect.TypeOf(fruits))

### compare != ถ้าข้างใน {} ค่าไม่ตรงกันก็จะไม่เป็นจริง

    fmt.Println([3]int{1, 2, 3} == [3]int{1, 3, 2})

### array used Copy by Value

array คือ declaration array

    newFruits := fruits
    newFruits[0] = "watermelon"

    fmt.Println(fruits)
    fmt.Println(newFruits)

### slices used copy by reference (pointer)

slice คือ การจองพื้นที่เอาไว้ก่อน สำหรับสร้าง array

    pointFruits := &fruits
    pointFruits[0] = "john"

    fmt.Println(fruits)
    fmt.Println(*pointFruits)

### array 2D

    a := [2][3]int{{1, 2, 3}, {3, 4, 5}}
    fmt.Println(a)

## slices

องประกอบของ​ slices

- data pointer ไปยังจุดเริ่มต้น
- len
- cap จุดท้ายสุด

fruits := []string{
"0",
"1",
"2",
"3",
"4",
}
fmt.Println(fruits[1:4])
fmt.Println(len(fruits[1:4])) // len = 3
fmt.Println(cap(fruits[1:4])) // cap = 4
// display : [1 2 3]
// เริ่มนับ 1 2 3 หยุดที่ 4

### slices copy by reference (pointer)

myfavorite := []string{
"apple",
"banana",
"papaya",
"orange",
}

yourfavorite := myfavorite
yourfavorite[0] = "guava"
fmt.Println(myfavorite[0]) // guava

### make a slices

s := make([]int, 5, 10)
fmt.Println("len", len(s), "cap", cap(s))

### Appends items to slices

x := [10]int{}
a := x[0:5]
b := x[5:7]
c := []int{1, 3, 6, 7, 8}

for i := range a {
a[i] = i \* i
}

b[0] = 98
b[1] = 99

fmt.Println("a", "len", len(a), "cap", cap(a), a)
fmt.Println("b", "len", len(b), "cap", cap(b), b)

// a = append(a, b[0], b[1]) b... คือการแตก array ให้เป็น element

### การใช้ append คือการเชื่อม array a กับ b ให้ต่อกัน เป็นของ slice a

a = append(a, b...)

//การเพิ่ม value เข้าไปใน slice a
a[len(a)-1] = 25
//display : a [0 1 4 9 16 98 25]

// เพิ่ม value ให้เต็ม slice a แต่ไม่เกิน cap
a = append(a, c...)

fmt.Println("a", a)
fmt.Println("b", b)
fmt.Println("x", x)

### append overcap จะเกิดการสร้าง new allocation memory ใหม่แบบซ่อนรูป unname \* ตามภาพ

#### append string + string

slice := append([]byte("hello "), "world"...)
fmt.Println(string(slice))

## map
