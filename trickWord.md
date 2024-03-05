# parameter & argument

parameter = คือใช้เรียกตอนประกาศฟังชั่น

สิ่งที่เรา pass ค่าเข้าไปเรียกว่า argument

## check the port on mac

-ข้อควรรู้ port protocal มันจะซ้ำกันไม่ได้

คำสั่งหา port
nc -l 8080 &

On macOS Big Sur
command : lsof -i -P | grep LISTEN | grep :$PORT
command : lsof -i -P | grep :8080

## kill the port

command : kill -9 PID
E.g. kill -9 60401
