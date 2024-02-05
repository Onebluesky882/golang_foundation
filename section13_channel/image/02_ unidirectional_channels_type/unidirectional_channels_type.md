# Unidirectional channels type

( Unidirectional channel types)
channel ที่ทำหน้าที่ได้อย่างเดียว ส่งได้อย่างเดียว หรือรับได้อย่างเดียว

Send only

chan<-

func f(ch chan<- int){
...
}

Receive only

<-chan

func f(ch <-chan int){
...
}

## Unidirectional channel type

จะไม่สามารถรับ chanal แบบ synchronous ได้
