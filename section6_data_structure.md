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

    newFruits := fruits
    newFruits[0] = "watermelon"

    fmt.Println(fruits)
    fmt.Println(newFruits)

### copy by reference

    pointFruits := &fruits
    pointFruits[0] = "john"

    fmt.Println(fruits)
    fmt.Println(*pointFruits)

### array 2D

    a := [2][3]int{{1, 2, 3}, {3, 4, 5}}
    fmt.Println(a)
