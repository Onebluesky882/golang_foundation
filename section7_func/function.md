# function

## function pattern

ประกาศ parameter และ return ต้อง type เดียวกัน

func name(parameter-list) (result-list){
return
}

## normal func

func avg(a, b float64) (c float64) {
c = (a + b) / 2
return c
}

func main() {
countDown(5)
}

## recursive function

ฟังชั่นที่เรียกทำงาน ด้วยตัวมันเอง จนกว่าเงื่อนไขจะสมบูรณ์
func countDown(c int) {
fmt.Println(c)
if c == 0 {
return
}
countDown(c - 1)
}
