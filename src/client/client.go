package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	log.SetFlags(log.Ltime)
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// copy conn --> os.stdout
	fmt.Println("connented to server successfully")
	go io.Copy(os.Stdout, conn)

	io.Copy(conn, os.Stdin)
}
