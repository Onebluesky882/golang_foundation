package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	log.SetFlags(log.Ltime)
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		log.Println("new client is connected!", conn.RemoteAddr())
		if err != nil {
			log.Fatal(err)
			return
		}
		// logic
		go countingDownHandler(conn)
	}

}

func countingDownHandler(conn net.Conn) {
	defer func() {
		io.WriteString(conn, "finish program will be close...\n")

		conn.Close()
	}()

	io.WriteString(conn, "Enter number : ")

	input := bufio.NewScanner(conn)

	count := scan(input)

	for {
		io.WriteString(conn, strconv.Itoa(count)+"\n")
		time.Sleep(time.Second)
		count--
		if count < 0 {
			io.WriteString(conn, "Enter number : ")
			count = scan(input)
		}

	}

}

func scan(input *bufio.Scanner) int {

	if ok := input.Scan(); !ok {
		log.Println("cannot scan value from conn")
		return 0
	}
	count, err := strconv.Atoi(input.Text())
	if err != nil {
		log.Println("can't convent value from text to int ", input.Text())
		return 0
	}
	return count
}
