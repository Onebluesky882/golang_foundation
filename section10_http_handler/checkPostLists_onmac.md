# command

-ข้อควรรู้ port protocal มันจะซ้ำกันไม่ได้

## check the port on mac

On macOS Big Sur
command : lsof -i -P | grep LISTEN | grep :$PORT

## kill the port

command : kill -9 PID
E.g. kill -9 60401
